package http

import (
	"net/http"

	"github.com/bitcoin-sv/paymail/service"
	"github.com/labstack/echo/v4"
)

// profileHandler is an http handler that supports BIP-270 requests.
type profileHandler struct {
	svc service.Profile
}

// NewP2PdDestHandler will create and return a new capabilitiesHandler.
func NewProfileHandler(svc service.Profile) *profileHandler {
	return &profileHandler{
		svc: svc,
	}
}

// RegisterRoutes will setup all routes with an echo group.
func (h *profileHandler) RegisterRoutes(g *echo.Group) {
	g.GET("/profile/:paymail", h.profileCreate)
}

// profileCreate generates a response object by forwarding the paymail to the pkiReader.
func (h *profileHandler) profileCreate(e echo.Context) error {
	resp := h.svc.ProfileReader(e.Request().Context(), e.Param("paymail"))
	return e.JSON(http.StatusAccepted, resp)
}
