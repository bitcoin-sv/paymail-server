package web

import (
	"net/http"

	"github.com/bitcoinschema/go-bitcoin"
	"github.com/labstack/echo"
	"github.com/nch-bowstave/paymail"
	"github.com/pkg/errors"
)

type PaymailService struct {
	svc AccountService
}

// NewBsvAlias will create a new bsvalias transport.
func NewBsvAlias(svc paymail.AccountService) *bsvalias {
	return &bsvalias{svc: svc}
}

// RegisterRoutes will setup the routes with the echo group.
func (b *bsvalias) RegisterRoutes(g *echo.Group) {
	g.POST(routePaymentDestination, b.PaymentDestination)
	g.GET(routePublicProfile, b.PublicProfile)
	g.GET(routeVerify, b.Verify)
}

// PaymentDestination obtains the preferred output script for a future transaction.
// Ref: https://docs.moneybutton.com/docs/paymail/paymail-04-payment-addressing.html
func (b *bsvalias) PaymentDestination(e echo.Context) error {
	handle := paymail.Handle(e.Param("handle"))
	senderRequest := new(paymail.SenderRequest)
	if err := e.Bind(senderRequest); err != nil {
		return e.JSON(http.StatusBadRequest, nil)
	}

	account, err := b.svc.Account(e.Request().Context(), handle)

	if err != nil {
		return errors.WithStack(err)
	}
	if account == nil {
		return e.JSON(http.StatusNotFound, nil)
	}

	output := &paymail.PaymentOutput{}
	if output.Output, err = bitcoin.ScriptFromAddress(account.Address); err != nil {
		return errors.WithStack(err)
	}

	return e.JSON(http.StatusOK, output)
}

// PaymentDestinationResponse
func (b *bsvalias) PaymentDestinationResponse(e echo.Context) error {
	return nil
}

// PublicProfile returns a users name and avatar url.
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
		AvatarURL: account.AvatarURL,
		Name:      account.Name,
	})
}

// Verify returns a users handle, pubkey, and boolean match (if successful).
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
		return e.JSON(http.StatusNotFound, nil)
	}
	return e.JSON(http.StatusOK, verification)
}
