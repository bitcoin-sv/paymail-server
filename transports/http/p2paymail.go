package http

import (
	"net/http"

	"github.com/bitcoin-sv/paymail/service"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// p2pDestHandler is an http handler that supports BIP-270 requests.
type p2PaymailHandler struct {
	svc service.P2Paymail
}

// NewP2PdDestHandler will create and return a new capabilitiesHandler.
func NewP2PaymailHandler(svc service.P2Paymail) *p2PaymailHandler {
	return &p2PaymailHandler{
		svc: svc,
	}
}

// RegisterRoutes will setup all routes with an echo group.
func (h *p2PaymailHandler) RegisterRoutes(g *echo.Group) {
	g.POST("/p2pDest/:paymail", h.p2pDest)
	g.POST("/p2pRawTx", h.p2pTxSubmit)
}

// p2pDest generates a response object using the static capabilities file.
func (h *p2PaymailHandler) p2pDest(e echo.Context) error {
	var args service.DestArgs
	if err := e.Bind(&args); err != nil {
		return errors.WithStack(err)
	}
	resp, err := h.svc.Destinations(e.Request().Context(), e.Param("paymail"), args)
	if err != nil {
		return errors.WithStack(err)
	}
	return e.JSON(http.StatusOK, resp)
}

// p2pTxSubmit sends the rawtx to payd which sends to chain
func (h *p2PaymailHandler) p2pTxSubmit(e echo.Context) error {
	var args service.TxSubmitArgs
	if err := e.Bind(&args); err != nil {
		return errors.WithStack(err)
	}
	resp, err := h.svc.RawTx(e.Request().Context(), args)
	if err != nil {
		return errors.WithStack(err)
	}
	return e.JSON(http.StatusOK, resp)
}
