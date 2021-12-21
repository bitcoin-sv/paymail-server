package service

import (
	"context"

	"github.com/libsv/p4-server/log"
	"github.com/nch-bowstave/paymail/data/payd"
)

type pki struct {
	l    log.Logger
	payd *payd.Payd
}

// NewPaymail will create and return a new paymail service.
func NewPki(l log.Logger, payd *payd.Payd) *pki {
	return &pki{
		l:    l,
		payd: payd,
	}
}

type PkiResponse struct {
	BsvAlias  string `json:"bsvalias"`
	Handle    string `json:"handle"`
	PublicKey string `json:"pubkey"`
}

// Paymail contains the handlers for paymail service endpoints.
type Pki interface {
	PkiReader(ctx context.Context, paymail string) (*PkiResponse, error)
}

func (svc *pki) PkiReader(ctx context.Context, paymail string) (*PkiResponse, error) {
	// TODO grab some pubkey from an account based on this paymail
	user, err := svc.payd.Owner(ctx)
	if err != nil {
		return nil, err
	}
	pki := &PkiResponse{
		BsvAlias:  "1.0",
		Handle:    paymail,
		PublicKey: user.Name,
	}
	return pki, nil
}
