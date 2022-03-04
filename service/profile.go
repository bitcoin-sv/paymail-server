package service

import (
	"context"

	"github.com/libsv/dpp-proxy/log"
	"github.com/nch-bowstave/paymail/data/payd"
	"github.com/nch-bowstave/paymail/data/sqlite"
)

type profile struct {
	l    log.Logger
	payd *payd.Payd
	str  sqlite.AliasStore
}

// NewPaymail will create and return a new paymail service.
func NewProfile(l log.Logger, payd *payd.Payd, str sqlite.AliasStore) *profile {
	return &profile{
		l:    l,
		payd: payd,
		str:  str,
	}
}

// ProfileResponse is the response object returned from a profile req.
type ProfileResponse struct {
	Name     string `json:"name,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
	ErrorMsg string `json:"error,omitempty"`
}

// Paymail contains the handlers for paymail service endpoints.
type Profile interface {
	ProfileReader(ctx context.Context, paymail string) *ProfileResponse
}

func (svc *profile) ProfileReader(ctx context.Context, paymail string) *ProfileResponse {
	errMsg := &ProfileResponse{
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

	profile := &ProfileResponse{
		Name:   user.Name,
		Avatar: user.AvatarURL,
	}
	return profile
}
