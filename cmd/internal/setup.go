package internal

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/bitcoin-sv/paymail/config"
	"github.com/bitcoin-sv/paymail/data"
	dppData "github.com/bitcoin-sv/paymail/data/dpp"
	"github.com/bitcoin-sv/paymail/data/payd"
	sql "github.com/bitcoin-sv/paymail/data/sqlite"
	"github.com/bitcoin-sv/paymail/docs"
	"github.com/bitcoin-sv/paymail/log"
	"github.com/bitcoin-sv/paymail/service"
	paymailHandlers "github.com/bitcoin-sv/paymail/transports/http"
	paymailMiddleware "github.com/bitcoin-sv/paymail/transports/http/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// Deps holds all the dependencies.
type Deps struct {
	PaymailService   service.Paymail
	PkiService       service.Pki
	ProfileService   service.Profile
	AliasService     service.Alias
	P2PaymailService service.P2Paymail
}

// SetupDeps will setup all required dependent services.
func SetupDeps(cfg config.Config, l log.Logger, db *sqlx.DB) *Deps {
	httpClient := &http.Client{Timeout: 5 * time.Second}
	if !cfg.PayD.Secure { // for testing, don't validate server cert
		// #nosec
		httpClient.Transport = &http.Transport{
			// #nosec
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	// stores
	httpDataClient := data.NewClient(httpClient)
	paydStore := payd.NewPayD(cfg.PayD, httpDataClient)
	dppClient := dppData.NewDPP(cfg.DPP, httpDataClient)
	sqlLiteStore := sql.NewSQLiteStore(db)

	// services
	paymailSvc := service.NewPaymail(l)
	pkiSvc := service.NewPki(l, paydStore, sqlLiteStore)
	profileSvc := service.NewProfile(l, paydStore, sqlLiteStore)
	aliasSvc := service.NewAlias(l, paydStore, sqlLiteStore)
	p2paymailSvc := service.NewP2Paymail(l, paydStore, dppClient, sqlLiteStore)

	return &Deps{
		PaymailService:   paymailSvc,
		PkiService:       pkiSvc,
		ProfileService:   profileSvc,
		AliasService:     aliasSvc,
		P2PaymailService: p2paymailSvc,
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
	c := e.Group("/")
	g := e.Group("/api")
	// handlers
	paymailHandlers.NewCapabilitiesHandler(deps.PaymailService).RegisterRoutes(c)
	paymailHandlers.NewPkiHandler(deps.PkiService).RegisterRoutes(g)
	paymailHandlers.NewProfileHandler(deps.ProfileService).RegisterRoutes(g)
	paymailHandlers.NewAliasHandler(deps.AliasService).RegisterRoutes(g)
	paymailHandlers.NewP2PaymailHandler(deps.P2PaymailService).RegisterRoutes(g)
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
