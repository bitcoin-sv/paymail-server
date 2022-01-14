package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nch-bowstave/paymail/models"
	"github.com/nch-bowstave/paymail/service"
	"github.com/pkg/errors"
)

// aliasHandler is an http handler that supports BIP-270 requests.
type aliasHandler struct {
	svc service.Alias
}

// NewP2PdDestHandler will create and return a new capabilitiesHandler.
func NewAliasHandler(svc service.Alias) *aliasHandler {
	return &aliasHandler{
		svc: svc,
	}
}

// RegisterRoutes will setup all routes with an echo group.
func (h *aliasHandler) RegisterRoutes(g *echo.Group) {
	g.POST("/alias", h.aliasCreate)
}

// aliasCreate generates a response object by forwarding the paymail to the aliasReader.
func (h *aliasHandler) aliasCreate(e echo.Context) error {
	var args models.AliasDetails
	if err := e.Bind(&args); err != nil {
		return errors.WithStack(err)
	}
	resp := h.svc.CreateAlias(e.Request().Context(), &args)
	return e.JSON(http.StatusAccepted, resp)
}
