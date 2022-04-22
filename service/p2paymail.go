package service

import (
	"context"
	"fmt"

	"github.com/libsv/dpp-proxy/log"
	"github.com/libsv/go-bc/spv"
	"github.com/libsv/go-dpp"
	dppData "github.com/nch-bowstave/paymail/data/dpp"
	paydData "github.com/nch-bowstave/paymail/data/payd"
	"github.com/nch-bowstave/paymail/data/sqlite"
	"github.com/nch-bowstave/paymail/models"
	"github.com/pkg/errors"
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
	dpp  dppData.DPP
	pki  sqlite.AliasStore
}

// NewPaymail will create and return a new paymail service.
func NewP2Paymail(l log.Logger, payd *paydData.Payd, dppClient dppData.DPP, pkiStr sqlite.AliasStore) *p2Paymail {
	return &p2Paymail{
		l:    l,
		payd: payd,
		dpp:  dppClient,
		pki:  pkiStr,
	}
}

// Paymail contains the handlers for paymail service endpoints.
type P2Paymail interface {
	Destinations(ctx context.Context, paymail string, args DestArgs) (*DestResponse, error)
	RawTx(ctx context.Context, args TxSubmitArgs) (*TxReceipt, error)
}

func (svc *p2Paymail) Destinations(ctx context.Context, paymail string, args DestArgs) (*DestResponse, error) {
	userID, err := svc.pki.GetUserID(ctx, paymail)
	if err != nil {
		return nil, err
	}
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

	destReq := models.PayRequest{
		PayToURL: fmt.Sprintf("http://%s/api/v1/payment/%s", svc.dpp.Host(), invoice.ID),
	}
	// grab some destinations from DPP
	response, err := svc.dpp.PaymentRequest(ctx, destReq)
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
	pcArgs := dpp.PaymentCreateArgs{PaymentID: args.Reference}
	invoice, err := svc.payd.GetInvoiceByID(ctx, args.Reference)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("invoice either doesn't exist or has been deleted %s", args.Reference))
	}
	if invoice.State == models.StateInvoiceDeleted {
		return nil, errors.New(fmt.Sprintf("invoice either doesn't exist or has been deleted %s", args.Reference))
	}
	req := dpp.Payment{
		MerchantData: dpp.Merchant{
			Name: args.MetaData.Signature,
			ExtendedData: map[string]interface{}{
				"paymail":   args.MetaData.Sender,
				"pki":       args.MetaData.PublicKey,
				"signature": args.MetaData.Signature,
			},
		},
		RawTX: &args.RawTx,
		Memo:  args.MetaData.Note,
		SPVEnvelope: &spv.Envelope{
			RawTx: args.RawTx,
		},
	}

	// TODO storing the requests for future reference - debugging
	// TODO check incoming signature across the TxID (use go-paymail)

	receipt, err := svc.payd.PaymentCreate(ctx, pcArgs, req)
	if err != nil {
		return nil, err
	}

	dest := &TxReceipt{
		TxID: receipt.TxID,
		Note: receipt.Memo,
	}
	return dest, nil
}
