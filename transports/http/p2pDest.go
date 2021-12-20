package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nch-bowstave/paymail/service"
	"github.com/pkg/errors"
)

// p2pDestHandler is an http handler that supports BIP-270 requests.
type p2pDestHandler struct {
	svc service.Paymail
}

// NewP2PdDestHandler will create and return a new capabilitiesHandler.
func NewP2PDestHandler(svc service.Paymail) *p2pDestHandler {
	return &p2pDestHandler{
		svc: svc,
	}
}

// RegisterRoutes will setup all routes with an echo group.
func (h *p2pDestHandler) RegisterRoutes(g *echo.Group) {
	g.GET("api/p2pDest/{alias}@{domain.tld}", h.p2pDestCreate)
}

// p2pDestCreate generates a response object using the static capabilities file.
func (h *p2pDestHandler) p2pDestCreate(e echo.Context) error {
	resp, err := h.svc.Capabilities(e.Request().Context())
	if err != nil {
		return errors.WithStack(err)
	}
	return e.JSON(http.StatusCreated, resp)
}
