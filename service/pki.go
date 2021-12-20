package service

import (
	"context"

	"github.com/libsv/p4-server/log"
)

type pki struct {
	l log.Logger
}

// NewPaymail will create and return a new paymail service.
func NewPki(l log.Logger) *pki {
	return &pki{
		l: l,
	}
}

type PkiResponse struct {
	BsvAlias  string `json:"bsvalias"`
	Handle    string `json:"handle"`
	PublicKey string `json:"pubkey"`
}

// Paymail contains the handlers for paymail service endpoints.
type Pki interface {
	PkiCreate(ctx context.Context, paymail string) (*PkiResponse, error)
}

func (svc *pki) PkiCreate(ctx context.Context, paymail string) (*PkiResponse, error) {
	// TODO grab some pubkey from an account.
	pki := &PkiResponse{
		BsvAlias:  "1.0",
		Handle:    paymail,
		PublicKey: "0203654654654798798465465413546546546546854654645",
	}
	// err = json.Unmarshal(data, pki)
	// if err != nil {
	// 	return nil, err
	// }
	return pki, nil
}
