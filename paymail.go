package paymail

import (
	"context"
	"errors"
	"net/mail"
	"time"
)

type (
	// Alias is the users alias
	Alias string

	//Handle
	Handle string

	// AccountArgs is the account creation fields.
	AccountArgs struct {
		Alias     string `json:"alias" db:"alias"`
		Name      string `json:"name" db:"name"`
		Email     string `json:"email" db:"email"`
		Mobile    string `json:"mobile" db:"mobile"`
		AvatarUrl string `json:"avatar_url" db:"avatar_url"`
	}

	// Account is the structure for a paymail account.
	Account struct {
		Alias      string `json:"alias" db:"alias"`
		Name       string `json:"name" db:"name"`
		Email      string `json:"email" db:"email"`
		Mobile     string `json:"mobile" db:"mobile"`
		AvatarUrl  string `json:"avatar_url" db:"avatar_url"`
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
		AvatarUrl string `json:"avatar_url" db:"avatar_url"`
		PublicKey string `json:"pubKey" db:"public_key"`
		Address   string `json:"address" db:"address"`
	}

	//PublicProfile is the public profile endpoint for a user.
	PublicProfile struct {
		AvatarUrl string `json:"avatar_url" db:"avatar_url"`
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

	// SenderRequest
	SenderRequest struct {
		Name      string    `json:"senderName"`
		Handle    string    `json:"senderHandle"`
		Date      time.Time `json:"dt"`
		Amount    int64     `json:"amount"`
		Purpose   string    `json:"purpose"`
		Signature string    `json:"signature"`
	}

	// PKI Public Key Infrastructure.
	PKI struct {
		BsvAlias  string `json:"bsvalias"`
		Handle    string `json:"handle"`
		PublicKey string `json:"pubkey"`
	}

	// P2PArgs is the payment request args.
	P2PArgs struct {
		Hex       string `json:"hex"`
		MetaData  `json:"metadata"`
		Reference string `json:"reference"`
	}

	// MetaData
	MetaData struct {
		Sender    string `json:"sender"`
		PublicKey string `json:"pubkey"`
		Signature string `json:"signature"`
		Note      string `json:"note"`
	}
	PaymentDestination struct {
		Outputs   []*PaymentOutput `json:"outputs"`   // A list of outputs
		Reference string           `json:"reference"` // A reference for the payment, created by the receiver of the transaction
	}

	PaymentOutput struct {
		Output string `json:"output"` // Hex encoded locking script
	}

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

	Domain interface {
		Verify(handle Handle) bool
	}
)

func (a *AccountArgs) Validate() error {
	_, err := mail.ParseAddress(string(a.Email))
	if err != nil {
		return errors.New("invalid email address")
	}
	return nil
}
