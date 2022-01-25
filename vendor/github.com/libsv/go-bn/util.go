package bn

import (
	"context"

	"github.com/libsv/go-bk/wif"
	"github.com/libsv/go-bn/models"
)

// UtilClient interfaces interaction with the util sub commands on a bitcoin node.
type UtilClient interface {
	ClearInvalidTransactions(ctx context.Context) (uint64, error)
	CreateMultiSig(ctx context.Context, n int, keys ...string) (*models.MultiSig, error)
	ValidateAddress(ctx context.Context, address string) (*models.ValidateAddress, error)
	SignMessageWithPrivKey(ctx context.Context, w *wif.WIF, msg string) (string, error)
	VerifySignedMessage(ctx context.Context, w *wif.WIF, signature, message string) (bool, error)
}

// NewUtilClient returns a client only capable of interfacing with the util sub commands on a bitcoin node.
func NewUtilClient(oo ...BitcoinClientOptFunc) UtilClient {
	return NewNodeClient(oo...)
}

// TODO: should not be cached
func (c *client) ClearInvalidTransactions(ctx context.Context) (uint64, error) {
	var resp uint64
	return resp, c.rpc.Do(ctx, "clearinvalidtransactions", &resp)
}

func (c *client) CreateMultiSig(ctx context.Context, n int, keys ...string) (*models.MultiSig, error) {
	var resp models.MultiSig
	return &resp, c.rpc.Do(ctx, "createmultisig", &resp, n, keys)
}

func (c *client) SignMessageWithPrivKey(ctx context.Context, w *wif.WIF, msg string) (string, error) {
	var resp string
	return resp, c.rpc.Do(ctx, "signmessagewithprivkey", &resp, w.String(), msg)
}

func (c *client) ValidateAddress(ctx context.Context, address string) (*models.ValidateAddress, error) {
	var resp models.ValidateAddress
	return &resp, c.rpc.Do(ctx, "validateaddress", &resp, address)
}

// TODO: Work out how to use
func (c *client) VerifySignedMessage(ctx context.Context, w *wif.WIF, signature, message string) (bool, error) {
	var resp bool
	return resp, c.rpc.Do(ctx, "verifymessage", &resp, w.String(), signature, message)
}
