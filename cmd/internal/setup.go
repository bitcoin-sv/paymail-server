package internal

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	p4soc "github.com/libsv/p4-server/transports/sockets"
	"github.com/nch-bowstave/paymail/config"
	"github.com/nch-bowstave/paymail/data"
	"github.com/nch-bowstave/paymail/data/payd"
	"github.com/nch-bowstave/paymail/data/sockets"
	"github.com/nch-bowstave/paymail/docs"
	"github.com/nch-bowstave/paymail/log"
	"github.com/nch-bowstave/paymail/service"
	paymailHandlers "github.com/nch-bowstave/paymail/transports/http"
	paymailMiddleware "github.com/nch-bowstave/paymail/transports/http/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
	smw "github.com/theflyingcodr/sockets/middleware"
	"github.com/theflyingcodr/sockets/server"

	"github.com/libsv/go-p4"
	"github.com/libsv/p4-server/data/noop"
	socData "github.com/libsv/p4-server/data/sockets"
)

// Deps holds all the dependencies.
type Deps struct {
	PaymailService        service.Paymail
	PaymentService        p4.PaymentService
	PaymentRequestService p4.PaymentRequestService
	ProofsService         p4.ProofsService
}

// SetupDeps will setup all required dependent services.
func SetupDeps(cfg config.Config, l log.Logger) *Deps {
	httpClient := &http.Client{Timeout: 5 * time.Second}
	if !cfg.PayD.Secure { // for testing, don't validate server cert
		// #nosec
		httpClient.Transport = &http.Transport{
			// #nosec
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	// stores
	paydStore := payd.NewPayD(cfg.PayD, data.NewClient(httpClient))

	// services
	paymailSvc := service.NewPaymail(l)
	paymentSvc := service.NewPayment(l, paydStore)
	paymentReqSvc := service.NewPaymentRequest(cfg.Server, paydStore, paydStore)
	if cfg.PayD.Noop {
		noopStore := noop.NewNoOp(log.Noop{})
		paymentSvc = service.NewPayment(log.Noop{}, noopStore)
		paymentReqSvc = service.NewPaymentRequest(cfg.Server, noopStore, noopStore)
	}
	proofService := service.NewProof(paydStore)

	return &Deps{
		PaymailService:        paymailSvc,
		PaymentService:        paymentSvc,
		PaymentRequestService: paymentReqSvc,
		ProofsService:         proofService,
	}
}

// SetupEcho will set up and return an echo server.
func SetupEcho(l log.Logger) *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.HTTPErrorHandler = paymailMiddleware.ErrorHandler(l)
	return e
}

// SetupSwagger will enable the swagger endpoints.
func SetupSwagger(cfg config.Server, e *echo.Echo) {
	docs.SwaggerInfo.Host = cfg.SwaggerHost
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}

// SetupHTTPEndpoints will register the http endpoints.
func SetupHTTPEndpoints(deps *Deps, e *echo.Echo) {
	g := e.Group("/")
	// handlers

	paymailHandlers.NewCapabilitiesHandler(deps.PaymailService).RegisterRoutes(g)

	// paymailHandlers.NewPaymentHandler(deps.PaymentService).RegisterRoutes(g)
	// paymailHandlers.NewPaymentRequestHandler(deps.PaymentRequestService).RegisterRoutes(g)
	// paymailHandlers.NewProofs(deps.ProofsService).RegisterRoutes(g)
}

// SetupSockets will setup handlers and socket server.
func SetupSockets(cfg config.Socket, e *echo.Echo) *server.SocketServer {
	g := e.Group("/")
	// create socket server
	s := server.New(
		server.WithMaxMessageSize(int64(cfg.MaxMessageBytes)),
		server.WithChannelTimeout(cfg.ChannelTimeout))

	// add middleware, with panic going first
	s.WithMiddleware(smw.PanicHandler, smw.Timeout(smw.NewTimeoutConfig()), smw.Metrics())

	p4soc.NewPaymentRequest().Register(s)
	p4soc.NewPayment().Register(s)
	paymailHandlers.NewProofs(service.NewProof(sockets.NewPayd(s))).RegisterRoutes(g)

	// this is our websocket endpoint, clients will hit this with the channelID they wish to connect to
	e.GET("/ws/:channelID", wsHandler(s))
	return s
}

// SetupHybrid will setup handlers for http=>socket communication.
func SetupHybrid(cfg config.Config, l log.Logger, e *echo.Echo) *server.SocketServer {
	g := e.Group("/")
	s := server.New(
		server.WithMaxMessageSize(int64(cfg.Sockets.MaxMessageBytes)),
		server.WithChannelTimeout(cfg.Sockets.ChannelTimeout))
	paymentStore := socData.NewPayd(s)
	paymentSvc := service.NewPayment(l, paymentStore)
	if cfg.PayD.Noop {
		noopStore := noop.NewNoOp(log.Noop{})
		paymentSvc = service.NewPayment(log.Noop{}, noopStore)
	}
	paymentReqSvc := service.NewPaymentRequestProxy(paymentStore, cfg.Transports, cfg.Server)
	proofsSvc := service.NewProof(paymentStore)

	paymailHandlers.NewPaymentHandler(paymentSvc).RegisterRoutes(g)
	paymailHandlers.NewPaymentRequestHandler(paymentReqSvc).RegisterRoutes(g)
	paymailHandlers.NewProofs(proofsSvc).RegisterRoutes(g)
	p4soc.NewHealthHandler().Register(s)

	e.GET("/ws/:channelID", wsHandler(s))
	return s
}

// wsHandler will upgrade connections to a websocket and then wait for messages.
func wsHandler(svr *server.SocketServer) echo.HandlerFunc {
	upgrader := websocket.Upgrader{}
	return func(c echo.Context) error {
		ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			return err
		}

		defer func() {
			_ = ws.Close()
		}()
		return svr.Listen(ws, c.Param("channelID"))
	}
}

// SetupSocketMetrics will setup the socket server metrics.
func SetupSocketMetrics(s *server.SocketServer) {
	// simple metrics
	gCo := promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "sockets",
		Subsystem: "server",
		Name:      "gauge_total_connections",
	})

	s.OnClientJoin(func(clientID, channelID string) {
		gCo.Inc()
	})

	s.OnClientLeave(func(clientID, channelID string) {
		gCo.Dec()
	})

	gCh := promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "sockets",
		Subsystem: "server",
		Name:      "gauge_total_channels",
	})

	s.OnChannelCreate(func(channelID string) {
		gCh.Inc()
	})

	s.OnChannelClose(func(channelID string) {
		gCh.Dec()
	})
}

// PrintDev outputs some useful dev information such as http routes
// and current settings being used.
func PrintDev(e *echo.Echo) {
	fmt.Println("==================================")
	fmt.Println("DEV mode, printing http routes:")
	for _, r := range e.Routes() {
		fmt.Printf("%s: %s\n", r.Method, r.Path)
	}
	fmt.Println("==================================")
	fmt.Println("DEV mode, printing settings:")
	for _, v := range viper.AllKeys() {
		fmt.Printf("%s: %v\n", v, viper.Get(v))
	}
	fmt.Println("==================================")
}
