package paymail

import "context"

// ref: https://docs.moneybutton.com/docs/paymail/paymail-03-public-key-infrastructure.html

type (
	// PKI Public Key Infrastructure.
	PKI struct {
		BsvAlias  string `json:"bsvalias"`
		Handle    string `json:"handle"`
		PublicKey string `json:"pubkey"`
	}
)

type (
	// PKIService contains the handlers for the PKI endpoints.
	PKIService interface {
		PKI(ctx context.Context, args Handle) (*PKI, error)
	}
)
