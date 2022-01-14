package sqlite

import (
	"context"
)

const (
	sqlCreateAlias = `
		INSERT INTO aliases(paymail, user_id)
		VALUES(:alias, :user_id)
	`

	sqlGetUserID = `
		SELECT user_id
		FROM aliases
		WHERE paymail = :alias
	`
)

type AliasStore interface {
	CreateAlias(ctx context.Context, alias string, userID uint64) error
	GetUserID(ctx context.Context, alias string) (uint64, error)
}

func (s *sqliteStore) CreateAlias(ctx context.Context, alias string, userID uint64) error {
	args := &struct {
		Alias  string `db:"alias"`
		UserID uint64 `db:"user_id"`
	}{
		Alias:  alias,
		UserID: userID,
	}
	_, err := s.db.NamedExec(sqlCreateAlias, args)
	if err != nil {
		return err
	}
	return nil
}

func (s *sqliteStore) GetUserID(ctx context.Context, alias string) (uint64, error) {
	dest := &struct {
		UserID uint64 `db:"user_id"`
	}{}
	err := s.db.GetContext(ctx, dest, sqlGetUserID, alias)
	if err != nil {
		return 0, err
	}
	return dest.UserID, nil
}
