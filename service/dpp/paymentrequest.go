package service

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	validator "github.com/theflyingcodr/govalidator"

	"github.com/bitcoin-sv/paymail/config"
	"github.com/libsv/go-dpp"
)

type paymentRequest struct {
	walletCfg   *config.Server
	destRdr     dpp.PaymentRequestReader
	merchantRdr dpp.Merchant
}

// NewPaymentRequest will setup and return a new PaymentRequest service that will generate outputs
// using the provided outputter which is defined in server config.
func NewPaymentRequest(walletCfg *config.Server, destRdr dpp.PaymentRequestReader, merchantRdr dpp.Merchant) *paymentRequest {
	return &paymentRequest{
		walletCfg:   walletCfg,
		destRdr:     destRdr,
		merchantRdr: merchantRdr,
	}
}

// PaymentRequest handles setting up a new PaymentRequest response and will validate that we have a paymentID.
func (p *paymentRequest) PaymentRequest(ctx context.Context, args dpp.PaymentRequestArgs) (*dpp.PaymentRequest, error) {
	if err := validator.New().
		Validate("paymentID", validator.NotEmpty(args.PaymentID)); err.Err() != nil {
		return nil, err
	}

	dests, err := p.destRdr.PaymentRequest(ctx, args)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get destinations for paymentID %s", args.PaymentID)
	}

	// get merchant information
	merchant := p.merchantRdr
	if merchant.ExtendedData == nil {
		merchant.ExtendedData = map[string]interface{}{}
	}
	// here we store paymentRef in extended data to allow some validation in payment flow
	merchant.ExtendedData["paymentReference"] = args.PaymentID
	return &dpp.PaymentRequest{
		Network:             dests.Network,
		SPVRequired:         dests.SPVRequired,
		Destinations:        dpp.PaymentDestinations{Outputs: dests.Destinations.Outputs},
		FeeRate:             dests.FeeRate,
		CreationTimestamp:   dests.CreationTimestamp,
		ExpirationTimestamp: dests.ExpirationTimestamp,
		PaymentURL:          fmt.Sprintf("http://%s/api/v1/payment/%s", p.walletCfg.FQDN, args.PaymentID),
		Memo:                fmt.Sprintf("invoice %s", args.PaymentID),
		MerchantData: &dpp.Merchant{
			AvatarURL:    merchant.AvatarURL,
			Name:         merchant.Name,
			Email:        merchant.Email,
			Address:      merchant.Address,
			ExtendedData: merchant.ExtendedData,
		},
	}, nil
}
