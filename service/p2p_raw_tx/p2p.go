package paymail

import "context"

// ref: https://docs.moneybutton.com/docs/paymail/paymail-06-p2p-transactions.html

type (
	// P2PArgs is the payment request args.
	P2PArgs struct {
		Hex       string `json:"hex"`
		MetaData  `json:"metadata"`
		Reference string `json:"reference"`
	}

	// MetaData contains extra information used in payment request args.
	MetaData struct {
		Sender    string `json:"sender"`
		PublicKey string `json:"pubkey"`
		Signature string `json:"signature"`
		Note      string `json:"note"`
	}

	// TransactionResponse is the txid resulting from the P2PArgs transaction.
	TransactionResponse struct {
		TxID string `json:"txid"`
		Note string `json:"note"`
	}
)

type (
	// P2PService contains the handlers for P2P endpoints.
	P2PService interface {
		RawTx(ctx context.Context, handle string, args P2PArgs) (*TransactionResponse, error)
	}
)
