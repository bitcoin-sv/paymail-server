package service

import (
	"context"

	"github.com/libsv/p4-server/log"
	"github.com/nch-bowstave/paymail/data/payd"
	"github.com/nch-bowstave/paymail/data/sqlite"
	"github.com/nch-bowstave/paymail/models"
	"github.com/pkg/errors"
)

type alias struct {
	l    log.Logger
	payd *payd.Payd
	str  sqlite.AliasStore
}

// NewAlias will create and return a new alias service.
func NewAlias(l log.Logger, payd *payd.Payd, str sqlite.AliasStore) *alias {
	return &alias{
		l:    l,
		payd: payd,
		str:  str,
	}
}

// Paymail contains the handlers for paymail service endpoints.
type Alias interface {
	CreateAlias(ctx context.Context, a *models.NewAliasDetails) *models.AliasResponse
}

func (svc *alias) CreateAlias(ctx context.Context, a *models.NewAliasDetails) *models.AliasResponse {
	user, err := svc.payd.CreateUser(ctx, models.UserDetails{
		Name:        a.Name,
		Email:       a.Email,
		Avatar:      a.Avatar,
		Address:     a.Address,
		PhoneNumber: a.PhoneNumber,
	})
	if err != nil {
		return &models.AliasResponse{
			Error: errors.Wrap(err, "User creation failed"),
		}
	}
	newAlias := &models.AliasResponse{
		UserID:  user.ID,
		Paymail: a.Paymail,
	}
	err = svc.str.CreateAlias(ctx, newAlias)
	if err != nil {
		return &models.AliasResponse{
			Error: errors.Wrap(err, "Alias creation failed"),
		}
	}
	return newAlias
}
