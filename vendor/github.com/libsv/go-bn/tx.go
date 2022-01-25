package bn

import (
	"context"

	imodels "github.com/libsv/go-bn/internal/models"
	"github.com/libsv/go-bn/models"
	"github.com/libsv/go-bt/v2"
)

// TransactionClient interfaces interaction with the transaction sub commands on a bitcoin node.
type TransactionClient interface {
	CreateRawTransaction(ctx context.Context, utxos bt.UTXOs, params models.ParamsCreateRawTransaction) (*bt.Tx, error)
	FundRawTransaction(ctx context.Context, tx *bt.Tx,
		opts *models.OptsFundRawTransaction) (*models.FundRawTransaction, error)
	RawTransaction(ctx context.Context, txID string) (*bt.Tx, error)
	SignRawTransaction(ctx context.Context, tx *bt.Tx,
		opts *models.OptsSignRawTransaction) (*models.SignedRawTransaction, error)
	SendRawTransaction(ctx context.Context, tx *bt.Tx, opts *models.OptsSendRawTransaction) (string, error)
	SendRawTransactions(ctx context.Context,
		params ...models.ParamsSendRawTransactions) (*models.SendRawTransactionsResponse, error)
}

// NewTransactionClient returns a client only capable of interfacing with the transaction sub commands
// on a bitcoin node.
func NewTransactionClient(oo ...BitcoinClientOptFunc) TransactionClient {
	return NewNodeClient(oo...)
}

func (c *client) CreateRawTransaction(ctx context.Context, utxos bt.UTXOs,
	params models.ParamsCreateRawTransaction) (*bt.Tx, error) {
	params.SetIsMainnet(c.isMainnet)
	var resp string
	if err := c.rpc.Do(ctx, "createrawtransaction", &resp, c.argsFor(&params, utxos.NodeJSON())...); err != nil {
		return nil, err
	}
	return bt.NewTxFromString(resp)
}

func (c *client) FundRawTransaction(ctx context.Context, tx *bt.Tx,
	opts *models.OptsFundRawTransaction) (*models.FundRawTransaction, error) {
	resp := imodels.InternalFundRawTransaction{FundRawTransaction: &models.FundRawTransaction{}}
	return resp.FundRawTransaction, c.rpc.Do(ctx, "fundrawtransaction", &resp, c.argsFor(opts, tx.String())...)
}

func (c *client) RawTransaction(ctx context.Context, txID string) (*bt.Tx, error) {
	var resp bt.Tx
	return &resp, c.rpc.Do(ctx, "getrawtransaction", &resp, txID, true)
}

func (c *client) SignRawTransaction(ctx context.Context, tx *bt.Tx,
	opts *models.OptsSignRawTransaction) (*models.SignedRawTransaction, error) {
	var resp imodels.InternalSignRawTransaction
	return resp.SignedRawTransaction, c.rpc.Do(ctx, "signrawtransaction", &resp, c.argsFor(opts, tx.String())...)
}

func (c *client) SendRawTransaction(ctx context.Context, tx *bt.Tx,
	opts *models.OptsSendRawTransaction) (string, error) {
	var resp string
	return resp, c.rpc.Do(ctx, "sendrawtransaction", &resp, c.argsFor(opts, tx.String())...)
}

func (c *client) SendRawTransactions(ctx context.Context,
	params ...models.ParamsSendRawTransactions) (*models.SendRawTransactionsResponse, error) {
	var resp models.SendRawTransactionsResponse
	return &resp, c.rpc.Do(ctx, "sendrawtransactions", &resp, params)
}
