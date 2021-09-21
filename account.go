package paymail

import (
	"context"
	"errors"
	"net/mail"
)

type (
	// AccountArgs is the account creation fields.
	AccountArgs struct {
		Alias     string `json:"alias" db:"alias"`
		Name      string `json:"name" db:"name"`
		Email     string `json:"email" db:"email"`
		Mobile    string `json:"mobile" db:"mobile"`
		AvatarURL string `json:"avatar_url" db:"avatar_url"`
	}

	// Account is the structure for a paymail account.
	Account struct {
		Alias      string `json:"alias" db:"alias"`
		Name       string `json:"name" db:"name"`
		Email      string `json:"email" db:"email"`
		Mobile     string `json:"mobile" db:"mobile"`
		AvatarURL  string `json:"avatar_url" db:"avatar_url"`
		PrivateKey string `db:"private_key"`
		PublicKey  string `json:"pubKey" db:"public_key"`
		Handle     string `json:"handle" db:"handle"`
		Address    string `json:"address" db:"address"`
	}

	// PublicAccount is the struct that contains only public viewable information.
	PublicAccount struct {
		Alias     string `json:"alias" db:"alias"`
		Handle    string `json:"handle" db:"handle"`
		Name      string `json:"name" db:"name"`
		AvatarURL string `json:"avatar_url" db:"avatar_url"`
		PublicKey string `json:"pubKey" db:"public_key"`
		Address   string `json:"address" db:"address"`
	}

	//PublicProfile is the public profile endpoint for a user.
	PublicProfile struct {
		AvatarURL string `json:"avatar_url" db:"avatar_url"`
		Name      string `json:"name" db:"name"`
	}

	// Verification is if the handle has correct matching pubkey.
	Verification struct {
		Handle string `json:"handle"`
		Pubkey string `json:"pubkey"`
		Match  bool   `json:"match"`
	}

	// VerificationArgs are the params to verify public key belongs to alias.
	VerificationArgs struct {
		Handle    string `json:"handle"`
		PublicKey string `json:"pubKey"`
	}
)

// Validate is used to check if email address is valid.
func (a *AccountArgs) Validate() error {
	_, err := mail.ParseAddress(a.Email)
	if err != nil {
		return errors.New("invalid email address")
	}
	return nil
}

type (
	// AccountService enforces validation of arguments and business rules.
	AccountService interface {
		Account(ctx context.Context, args Handle) (*PublicAccount, error)
		Create(ctx context.Context, req AccountArgs) error
		Verify(ctx context.Context, args VerificationArgs) (*Verification, error)
	}

	// AccountWriter creates a new paymail account.
	AccountWriter interface {
		Create(ctx context.Context, req Account) error
	}

	// AccountReader is used to retrieve paymail accounts.
	AccountReader interface {
		Account(ctx context.Context, args Handle) (*PublicAccount, error)
	}

	// AccountReaderWriter is the wrapper around AccountWriter and AccountReader.
	AccountReaderWriter interface {
		AccountReader
		AccountWriter
	}
)
