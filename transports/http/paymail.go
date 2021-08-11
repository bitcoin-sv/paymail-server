package web

import (
	"net/http"

	"github.com/bitcoinschema/go-bitcoin"
	"github.com/labstack/echo"
	"github.com/nch-bowstave/paymail"
	"github.com/pkg/errors"
)

type bsvalias struct {
	svc paymail.AccountService
}

// NewBsvAlias will create a new bsvalias transport.
func NewBsvAlias(svc paymail.AccountService) *bsvalias {
	return &bsvalias{svc: svc}
}

// RegisterRoutes will setup the routes with the echo group.
func (b *bsvalias) RegisterRoutes(g *echo.Group) {
	g.GET(routePki, b.PKI)
	g.POST(routePaymentDestination, b.PaymentDestination)
	g.GET(routePublicProfile, b.PublicProfile)
	g.GET(routeVerify, b.Verify)
}

func (b *bsvalias) PKI(e echo.Context) error {
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

func (b *bsvalias) PaymentDestination(e echo.Context) error {
	handle := paymail.Handle(e.Param("handle"))
	paymentRequest := new(paymail.PaymentRequest)
	if err := e.Bind(paymentRequest); err != nil {
		return e.JSON(http.StatusBadRequest, nil)
	}

	account, err := b.svc.Account(e.Request().Context(), handle)
	if err != nil {
		return errors.WithStack(err)
	}
	if account == nil {
		return e.JSON(http.StatusNotFound, nil)
	}

	output := &paymail.PaymentOutput{
		Address:  account.Address,
		Satoshis: paymentRequest.Satoshis,
	}

	if output.Script, err = bitcoin.ScriptFromAddress(account.Address); err != nil {
		return errors.WithStack(err)
	}

	return e.JSON(http.StatusOK, output)
}

func (b *bsvalias) PublicProfile(e echo.Context) error {
	handle := paymail.Handle(e.Param("handle"))
	account, err := b.svc.Account(e.Request().Context(), handle)
	if err != nil {
		return errors.WithStack(err)
	}
	if account == nil {
		return e.JSON(http.StatusNotFound, nil)
	}
	return e.JSON(http.StatusOK, &paymail.PublicProfile{
		AvatarUrl: account.AvatarUrl,
		Name:      account.Name,
	})
}

func (b *bsvalias) Verify(e echo.Context) error {
	handle := e.Param("handle")
	pubkey := e.Param("pubkey")
	verification, err := b.svc.Verify(e.Request().Context(), paymail.VerificationArgs{
		Handle:    handle,
		PublicKey: pubkey,
	})
	if err != nil {
		return errors.WithStack(err)
	}
	if verification == nil {
		e.JSON(http.StatusNotFound, nil)
	}
	return e.JSON(http.StatusOK, verification)
}
