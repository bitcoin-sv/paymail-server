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
	CreateAlias(ctx context.Context, a *models.AliasDetails) *models.AliasDetails
}

func (svc *alias) CreateAlias(ctx context.Context, a *models.AliasDetails) *models.AliasDetails {
	err := svc.str.CreateAlias(ctx, a)
	if err != nil {
		return &models.AliasDetails{
			Error: errors.Wrap(err, "alias creation failed"),
		}
	}
	return a
}
