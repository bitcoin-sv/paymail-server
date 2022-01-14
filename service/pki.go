package service

import (
	"context"
	"fmt"

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
	PublicKey string `json:"pubkey,omitempty"`
	ErrorMsg  string `json:"error,omitempty"`
}

// Paymail contains the handlers for paymail service endpoints.
type Pki interface {
	PkiReader(ctx context.Context, paymail string) *PkiResponse
}

func (svc *pki) PkiReader(ctx context.Context, paymail string) *PkiResponse {
	userID := GetUserIDFromPaymail(paymail)
	// domain = p[1]
	// TODO check domain matches one of our env paymail domains
	user, err := svc.payd.User(ctx, userID)
	if err != nil {
		return &PkiResponse{
			BsvAlias: "1.0",
			Handle:   paymail,
			ErrorMsg: "Not found at this domain.",
		}
	}

	pk := fmt.Sprintf("%s", user.ExtendedData["pki"])

	pki := &PkiResponse{
		BsvAlias:  "1.0",
		Handle:    paymail,
		PublicKey: pk,
	}
	return pki
}
