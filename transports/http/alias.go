package http

import (
	"net/http"

	"github.com/bitcoin-sv/paymail/models"
	"github.com/bitcoin-sv/paymail/service"
	"github.com/labstack/echo/v4"
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
	var args models.NewAliasDetails
	if err := e.Bind(&args); err != nil {
		return errors.WithStack(err)
	}
	resp, err := h.svc.CreateAlias(e.Request().Context(), &args)
	if err != nil {
		return e.JSON(http.StatusBadRequest, &models.AliasResponse{
			Error: err.Error(),
		})
	}
	return e.JSON(http.StatusAccepted, resp)
}
