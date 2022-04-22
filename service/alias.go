package service

import (
	"context"

	"github.com/bitcoin-sv/paymail/data/payd"
	"github.com/bitcoin-sv/paymail/data/sqlite"
	"github.com/bitcoin-sv/paymail/models"
	"github.com/libsv/dpp-proxy/log"
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
	CreateAlias(ctx context.Context, a *models.NewAliasDetails) (*models.AliasResponse, error)
}

func (svc *alias) CreateAlias(ctx context.Context, a *models.NewAliasDetails) (*models.AliasResponse, error) {
	if a.Paymail == "" {
		return nil, errors.New("must include paymail to assign")
	}
	user, err := svc.payd.CreateUser(ctx, models.UserDetails{
		Name:        a.Name,
		Email:       a.Email,
		Avatar:      a.Avatar,
		Address:     a.Address,
		PhoneNumber: a.PhoneNumber,
	})
	if err != nil {
		return nil, errors.Wrap(err, "user creation failed")
	}
	newAlias := &models.AliasResponse{
		UserID:  user.ID,
		Paymail: a.Paymail,
	}
	err = svc.str.CreateAlias(ctx, newAlias)
	if err != nil {
		return nil, errors.Wrap(err, "alias creation failed")
	}
	return newAlias, nil
}
