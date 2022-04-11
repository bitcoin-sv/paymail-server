package noop

import (
	"context"

	"github.com/libsv/dpp-proxy/log"
	"github.com/libsv/go-bt/v2"
	"github.com/libsv/go-bt/v2/bscript"

	"github.com/libsv/go-dpp"
)

type noop struct {
	l log.Logger
}

// NewNoOp will setup and return a new no operational data store for
// testing purposes. Useful if you want to explore endpoints without
// integrating with a wallet.
func NewNoOp(l log.Logger) *noop {
	l.Info("using NOOP data store")
	return &noop{}
}

// PaymentCreate will post a request to payd to validate and add the txos to the wallet.
//
// If invalid a non 204 status code is returned.
func (n *noop) PaymentCreate(ctx context.Context, args dpp.PaymentCreateArgs, req dpp.Payment) (*dpp.PaymentACK, error) {
	n.l.Info("hit noop.PaymentCreate")
	return &dpp.PaymentACK{}, nil
}

// Owner will return information regarding the owner of a payd wallet.
//
// In this example, the payd wallet has no auth, in proper implementations auth would
// be enabled and a cookie / oauth / bearer token etc would be passed down.
func (n *noop) User(ctx context.Context) (*dpp.Merchant, error) {
	n.l.Info("hit noop.Owner")
	return &dpp.Merchant{
		AvatarURL:    "noop",
		Name:         "noop",
		Email:        "noop",
		Address:      "noop",
		ExtendedData: nil,
	}, nil
}

func (n *noop) Destinations(ctx context.Context, args dpp.PaymentRequestArgs) (*dpp.Destinations, error) {
	n.l.Info("hit noop.Destinations")
	return &dpp.Destinations{
		Outputs: []dpp.Output{{
			Amount:        0,
			LockingScript: &bscript.Script{},
			Description:   "noop",
		}},
		Fees: bt.NewFeeQuote(),
	}, nil
}
