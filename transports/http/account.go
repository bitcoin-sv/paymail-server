package web

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/nch-bowstave/paymail"
	"github.com/pkg/errors"
)

type account struct {
	svc paymail.AccountService
}

// NewAccount will create a new paymail account transport.
func NewAccount(svc paymail.AccountService) *account {
	return &account{svc: svc}
}

// RegisterRoutes will setup the routes with the echo group.
func (a *account) RegisterRoutes(g *echo.Group) {
	g.POST(routeRegister, a.Account)
}

// Account will create a new paymail account if valid request.
func (a *account) Account(e echo.Context) error {
	var args paymail.AccountArgs
	if err := e.Bind(&args); err != nil {
		return err
	}
	if err := a.svc.Create(e.Request().Context(), args); err != nil {
		return errors.WithStack(err)
	}
	return e.JSON(http.StatusCreated, struct{}{})
}
