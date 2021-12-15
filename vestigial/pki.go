package web

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/nch-bowstave/paymail"
	"github.com/pkg/errors"
)

type pki struct {
	svc paymail.AccountService
}

// NewPKI will create a new pki transport.
func NewPKI(svc paymail.AccountService) *pki {
	return &pki{svc: svc}
}

// RegisterRoutes will setup the routes with the echo group.
func (b *pki) RegisterRoutes(g *echo.Group) {
	g.GET(routePki, b.PKI)
}

// PKI Public Key Infrastructure returns a users handle, bsv alias version, and public key for a user handle (if exists),
func (b *pki) PKI(e echo.Context) error {
	handle := paymail.Handle(e.Param("handle"))
	account, err := b.svc.Account(e.Request().Context(), handle)
	if err != nil {
		return errors.WithStack(err)
	}
	if account == nil {
		return e.JSON(http.StatusNotFound, nil)
	}
	return e.JSON(http.StatusOK, &paymail.PKI{
		BsvAlias:  "1.0",
		Handle:    account.Handle,
		PublicKey: account.PublicKey,
	})
}
