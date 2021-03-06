package sqlite

import (
	"context"
	"errors"

	"github.com/bitcoin-sv/paymail/models"
)

const (
	sqlCreateAlias = `
		INSERT INTO aliases(paymail, user_id)
		VALUES(:paymail, :user_id)
	`

	sqlGetUserID = `
		SELECT user_id
		FROM aliases
		WHERE paymail = :paymail
	`
)

type AliasStore interface {
	CreateAlias(ctx context.Context, a *models.AliasResponse) error
	GetUserID(ctx context.Context, paymail string) (uint64, error)
}

func (s *sqliteStore) CreateAlias(ctx context.Context, a *models.AliasResponse) error {
	if a.Paymail == "" || a.UserID == 0 {
		return errors.New("must include paymail and user_id to assign")
	}
	_, err := s.db.NamedExec(sqlCreateAlias, a)
	if err != nil {
		return err
	}
	return nil
}

func (s *sqliteStore) GetUserID(ctx context.Context, paymail string) (uint64, error) {
	dest := &struct {
		UserID uint64 `db:"user_id"`
	}{}
	err := s.db.GetContext(ctx, dest, sqlGetUserID, paymail)
	if err != nil {
		return 0, err
	}
	return dest.UserID, nil
}
