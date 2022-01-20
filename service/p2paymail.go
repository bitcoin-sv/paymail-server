package service

import (
	"context"
	"fmt"

	"github.com/libsv/go-bt/v2"
	"github.com/libsv/go-p4"
	"github.com/libsv/p4-server/log"
	data "github.com/nch-bowstave/paymail/data/p4"
	paydData "github.com/nch-bowstave/paymail/data/payd"
	"github.com/nch-bowstave/paymail/models"
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
	payd *paydData.Payd
	p4   data.P4
}

// GetUserIDFromPaymail
func GetUserIDFromPaymail(paymail string) uint64 {
	// TODO lookup the paymail in our database and return a user_id
	return 1
}

// NewPaymail will create and return a new paymail service.
func NewP2Paymail(l log.Logger, payd *paydData.Payd, p4Client data.P4) *p2Paymail {
	return &p2Paymail{
		l:    l,
		payd: payd,
		p4:   p4Client,
	}
}

// Paymail contains the handlers for paymail service endpoints.
type P2Paymail interface {
	Destinations(ctx context.Context, paymail string, args DestArgs) (*DestResponse, error)
	RawTx(ctx context.Context, args TxSubmitArgs) (*TxReceipt, error)
}

func (svc *p2Paymail) Destinations(ctx context.Context, paymail string, args DestArgs) (*DestResponse, error) {
	userID := GetUserIDFromPaymail(paymail)
	req := &models.InvoiceCreate{
		UserID:      userID,
		Satoshis:    args.Satoshis,
		SPVRequired: false,
	}
	// create an invoice on payd for the amount specified in the args
	invoice, err := svc.payd.CreateInvoice(ctx, req)
	if err != nil {
		return nil, err
	}

	destReq := models.P4PayRequest{
		PayToURL: fmt.Sprintf("http://%s/api/v1/payment/%s", svc.p4.Host(), invoice.ID),
	}
	// grab some destinations from P4
	response, err := svc.p4.PaymentRequest(ctx, destReq)
	if err != nil {
		return nil, err
	}

	var outputs []Output
	for _, output := range response.Destinations.Outputs {
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
	pcArgs := p4.PaymentCreateArgs{PaymentID: args.Reference}
	req := p4.Payment{
		MerchantData: p4.Merchant{
			Name: args.MetaData.Signature,
			ExtendedData: map[string]interface{}{
				"paymail":   args.MetaData.Sender,
				"pki":       args.MetaData.PublicKey,
				"signature": args.MetaData.Signature,
			},
		},
		RawTX: &args.RawTx,
		Memo:  args.MetaData.Note,
	}

	// TODO storing the requests for future reference - debugging
	// TODO check incoming signature across the TxID (use go-paymail)

	receipt, err := svc.payd.PaymentCreate(ctx, pcArgs, req)
	if err != nil {
		return nil, err
	}

	tx, err := bt.NewTxFromString(*receipt.Payment.RawTX)
	if err != nil {
		return nil, err
	}
	txid := tx.TxID()

	dest := &TxReceipt{
		TxID: txid,
		Note: receipt.Payment.Memo,
	}
	return dest, nil
}
