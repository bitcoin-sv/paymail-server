package service

import (
	"context"
	"fmt"

	"github.com/libsv/go-p4"
	"github.com/libsv/p4-server/log"
	paydMessages "github.com/libsv/payd"
	"github.com/nch-bowstave/paymail/data/payd"
)

// ref: https://docs.moneybutton.com/docs/paymail/paymail-06-p2p-transactions.html

// Dest is the p2p destination request args
type DestArgs struct {
	Satoshis uint64 `json:"satoshis"`
}

type DestResponse struct {
	Reference string   `json:"reference"`
	Outputs   []Output `json:"outputs"`
}

type Output struct {
	Script   string `json:"script"`
	Satoshis uint64 `json:"satoshis"`
}

// TxSubmitArgs is the p2p submit transaction args
type TxSubmitArgs struct {
	RawTx     string   `json:"hex"`
	Reference string   `json:"reference"`
	MetaData  MetaData `json:"metadata"`
}

type MetaData struct {
	Sender    string `json:"sender"`
	PublicKey string `json:"pubkey"`
	Signature string `json:"signature"`
	Note      string `json:"note"`
}

// Receipt is the txid resulting from the P2PArgs transaction.
type TxReceipt struct {
	TxID string `json:"txid"`
	Note string `json:"note"`
}

type p2Paymail struct {
	l    log.Logger
	payd *payd.Payd
}

// NewPaymail will create and return a new paymail service.
func NewP2Paymail(l log.Logger, payd *payd.Payd) *p2Paymail {
	return &p2Paymail{
		l:    l,
		payd: payd,
	}
}

// Paymail contains the handlers for paymail service endpoints.
type P2Paymail interface {
	Destinations(ctx context.Context, paymail string, args DestArgs) (*DestResponse, error)
	RawTx(ctx context.Context, args TxSubmitArgs) (*TxReceipt, error)
}

func (svc *p2Paymail) Destinations(ctx context.Context, paymail string, args DestArgs) (*DestResponse, error) {
	req := &paydMessages.InvoiceCreate{
		Satoshis:    args.Satoshis,
		SPVRequired: false,
	}
	// create an invoice on payd for the amount specified in the args
	invoice, err := svc.payd.CreateInvoice(ctx, req)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v", invoice)

	// grab some destinations from PayD
	response, err := svc.payd.Destinations(ctx, p4.PaymentRequestArgs{
		PaymentID: invoice.ID,
	})
	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v", response)

	var outputs []Output
	for _, output := range response.Outputs {
		outputs = append(outputs, Output{
			Script:   output.Script,
			Satoshis: output.Amount,
		})
	}
	dest := &DestResponse{
		Reference: invoice.ID,
		Outputs:   outputs,
	}
	return dest, nil
}

func (svc *p2Paymail) RawTx(ctx context.Context, args TxSubmitArgs) (*TxReceipt, error) {
	// TODO payment submit to PayD
	user, err := svc.payd.Owner(ctx)
	if err != nil {
		return nil, err
	}
	receipt := &TxReceipt{
		TxID: "685498798451651654654654654689749/874171089704897408740189410869406840",
		Note: user.Name,
	}
	return receipt, nil
}
