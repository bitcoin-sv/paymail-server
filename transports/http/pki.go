package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nch-bowstave/paymail/service"
	"github.com/pkg/errors"
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
	g.GET("api/pki/:paymail", h.pkiCreate)
}

// p2pDestCreate generates a response object using the static capabilities file.
func (h *pkiHandler) pkiCreate(e echo.Context) error {
	resp, err := h.svc.PkiCreate(e.Request().Context(), e.Param("paymail"))
	if err != nil {
		return errors.WithStack(err)
	}
	return e.JSON(http.StatusCreated, resp)
}
