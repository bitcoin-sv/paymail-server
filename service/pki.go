package service

import (
	"context"
	"crypto/elliptic"
	"encoding/hex"

	"github.com/bitcoinsv/bsvd/bsvec"
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
	// pki, err := svc.payd.Pki(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// replace this once pki is implemented within PayD
	key, err := bsvec.NewPrivateKey(elliptic.P256())
	if err != nil {
		return nil, err
	}

	pubkey := key.PubKey()

	pki := &PkiResponse{
		BsvAlias:  "1.0",
		Handle:    paymail,
		PublicKey: hex.EncodeToString(pubkey.SerializeCompressed()),
	}
	return pki, nil
}
