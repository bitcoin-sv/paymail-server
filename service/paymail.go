package service

import (
	"context"

	"github.com/bitcoinschema/go-bitcoin"
	"github.com/nch-bowstave/paymail"
	"github.com/pkg/errors"
)

type paymailService struct {
	domain string
	rw     paymail.AccountReaderWriter
}

// NewPaymailService
func NewPaymailService(rw paymail.AccountReaderWriter, domain string) *paymailService {
	return &paymailService{
		domain: domain,
		rw:     rw,
	}
}

// Create creates a new user account.
func (p *paymailService) Create(ctx context.Context, req paymail.AccountArgs) error {
	if err := req.Validate(); err != nil {
		return err
	}
	account := paymail.Account{
		Handle:    req.Alias + "@" + p.domain,
		Alias:     req.Alias,
		Name:      req.Name,
		AvatarURL: req.AvatarURL,
		Email:     req.Email,
		Mobile:    req.Mobile,
	}

	var err error
	if account.PrivateKey, err = bitcoin.CreatePrivateKeyString(); err != nil {
		return errors.Wrap(err, "error creating private key")
	}

	if account.PublicKey, err = bitcoin.PubKeyFromPrivateKeyString(account.PrivateKey, true); err != nil {
		return errors.Wrap(err, "error creating public key")
	}

	if account.Address, err = bitcoin.GetAddressFromPrivateKeyString(account.PrivateKey, true); err != nil {
		return errors.Wrap(err, "error creating bitcoin address")
	}

	if err = p.rw.Create(ctx, account); err != nil {
		return errors.Wrap(err, "failed to create account")
	}
	return nil
}

// Account returns an account for a user.
func (p *paymailService) Account(ctx context.Context, args paymail.Handle) (*paymail.PublicAccount, error) {
	return p.rw.Account(ctx, args)
}

// Verify looks up the users account and checks if pubkey matches user param.
func (p *paymailService) Verify(ctx context.Context, args paymail.VerificationArgs) (*paymail.Verification, error) {
	account, err := p.Account(ctx, paymail.Handle(args.Handle))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if account == nil {
		return nil, nil
	}
	return &paymail.Verification{
		Handle: args.Handle,
		Pubkey: args.PublicKey,
		Match:  args.PublicKey == account.PublicKey,
	}, nil
}
