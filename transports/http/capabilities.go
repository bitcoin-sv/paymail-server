package http

import (
	"net/http"

	"github.com/bitcoin-sv/paymail/service"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// capabilitiesHandler is the capabilities discovery document request handler.
type capabilitiesHandler struct {
	svc service.Paymail
}

// NewCapabilitiesHandler will create and return a new capabilitiesHandler.
func NewCapabilitiesHandler(svc service.Paymail) *capabilitiesHandler {
	return &capabilitiesHandler{
		svc: svc,
	}
}

// RegisterRoutes will setup all routes with an echo group.
func (h *capabilitiesHandler) RegisterRoutes(g *echo.Group) {
	g.GET(".well-known/bsvalias.json", h.capabilitiesResponse)
	g.GET(".well-known/bsvalias", h.capabilitiesResponse) // backward compatible version
}

// capabilitiesResponse generates a response object using the static capabilities file.
func (h *capabilitiesHandler) capabilitiesResponse(e echo.Context) error {
	data, err := h.svc.Capabilities(e.Request().Context(), e.Request().URL.Path)
	if err != nil {
		return errors.WithStack(err)
	}
	return e.JSONBlob(http.StatusOK, data)
}
