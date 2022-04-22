package http

import (
	"net/http"

	"github.com/bitcoin-sv/paymail/service"
	"github.com/labstack/echo/v4"
)

// pkiHandler is an http handler that supports BIP-270 requests.
type pkiHandler struct {
	svc service.Pki
}

// NewP2PdDestHandler will create and return a new capabilitiesHandler.
func NewPkiHandler(svc service.Pki) *pkiHandler {
	return &pkiHandler{
		svc: svc,
	}
}

// RegisterRoutes will setup all routes with an echo group.
func (h *pkiHandler) RegisterRoutes(g *echo.Group) {
	g.GET("/pki/:paymail", h.pkiCreate)
}

// pkiCreate generates a response object by forwarding the paymail to the pkiReader.
func (h *pkiHandler) pkiCreate(e echo.Context) error {
	resp := h.svc.PkiReader(e.Request().Context(), e.Param("paymail"))
	return e.JSON(http.StatusAccepted, resp)
}
