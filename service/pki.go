package service

import (
	"context"
	"fmt"

	"github.com/libsv/p4-server/log"
	"github.com/nch-bowstave/paymail/data/payd"
	"github.com/nch-bowstave/paymail/data/sqlite"
)

type pki struct {
	l    log.Logger
	payd *payd.Payd
	str  sqlite.AliasStore
}

// NewPaymail will create and return a new paymail service.
func NewPki(l log.Logger, payd *payd.Payd, str sqlite.AliasStore) *pki {
	return &pki{
		l:    l,
		payd: payd,
		str:  str,
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
	errMsg := &PkiResponse{
		BsvAlias: "1.0",
		Handle:   paymail,
		ErrorMsg: "Not found at this domain.",
	}

	userID, err := svc.str.GetUserID(ctx, paymail)
	if err != nil {
		return errMsg
	}

	user, err := svc.payd.User(ctx, userID)
	if err != nil {
		return errMsg
	}

	pki := &PkiResponse{
		BsvAlias:  "1.0",
		Handle:    paymail,
		PublicKey: fmt.Sprintf("%s", user.ExtendedData["pki"]),
	}
	return pki
}
