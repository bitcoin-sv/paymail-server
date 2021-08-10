package sql

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/nch-bowstave/paymail"
	"github.com/nch-bowstave/paymail/config"
)

const (
	insertAccount = `
	INSERT INTO accounts(alias, handle, name, avatar_url, private_key, public_key, address, email, mobile)
	VALUES(:alias, :handle, :name, :avatar_url, :private_key, :public_key, :address, :email, :mobile)
	`

	sqlAccount = `
	SELECT alias, name, handle, avatar_url, address, public_key
  FROM accounts
  WHERE handle = ?
	`
)

type paymailDb struct {
	dbType config.DbType
	db     *sqlx.DB
	sqls   map[config.DbType]map[string]string
}

// NewPaymailDb will setup and return a new paymail store.
func NewPaymailDb(db *sqlx.DB, dbType config.DbType) *paymailDb {
	return &paymailDb{
		dbType: dbType,
		db:     db,
		sqls: map[config.DbType]map[string]string{
			config.DBMySql: {
				insertAccount: insertAccount,
			},
			config.DBPostgres: {
				insertAccount: insertAccount,
				sqlAccount:    sqlAccount,
			},
		},
	}
}

func (h *paymailDb) Create(ctx context.Context, req paymail.Account) error {
	tx, err := h.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()
	if _, err := tx.NamedExecContext(ctx, h.sqls[h.dbType][insertAccount], req); err != nil {
		return errors.Wrap(err, "failed to create account")
	}
	return errors.Wrap(tx.Commit(), "failed to commit tx")
}

// Account will return an account by paymail address (handle).
func (h *paymailDb) Account(ctx context.Context, args paymail.Handle) (*paymail.PublicAccount, error) {
	var bh paymail.PublicAccount
	if err := h.db.GetContext(ctx, &bh, h.db.Rebind(sqlAccount), args); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("could not find account")
			// return nil, lathos.NewErrNotFound("N001", "could not find account")
		}
		return nil, errors.Wrapf(err, "failed to get account using %s", args)
	}
	return &bh, nil
}
