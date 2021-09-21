package paymail

import (
	"context"
	"time"
)

// ref: https://docs.moneybutton.com/docs/paymail/paymail-04-payment-addressing.html

type (
	// PaymentDestination contains a list of outputs and a reference created by the receiver of the transaction.
	PaymentDestination struct {
		Outputs   []*PaymentOutput `json:"outputs"`   // A list of outputs
		Reference string           `json:"reference"` // A reference for the payment, created by the receiver of the transaction
	}

	// PaymentOutput is the transaction locking script.
	PaymentOutput struct {
		Output string `json:"output"` // Hex encoded locking script
	}

	// SenderRequest is the structure for a payment request.
	SenderRequest struct {
		Name      string    `json:"senderName"`
		Handle    string    `json:"senderHandle"`
		Date      time.Time `json:"dt"` // time.Now().Format(time.RFC3339)
		Amount    int64     `json:"amount"`
		Purpose   string    `json:"purpose"`
		Signature string    `json:"signature"`
	}

	// ReceiverResponse is data to be used as part of callback in PaymentDestinationApproval.
	ReceiverResponse struct {
		Token     string `json:"token"`
		Signature string `json:"signature"`
	}

	// PaymentDestinationApproval is part of the callback from a payment destination request.
	PaymentDestinationApproval struct {
		Token     string `json:"token"`
		Signature string `json:"signature"`
		Output    string `json:"output"`
	}
)

type (
	// PaymentService contains the handlers for payment service endpoints.
	PaymentService interface {
		PaymentDestination(ctx context.Context, handle Handle) (*PaymentDestination, error)
		PaymentDestinationApproval(ctx context.Context, handle Handle) (*PaymentDestination, error)
		SenderRequest(ctx context.Context, handle Handle, args SenderRequest) (*PaymentOutput, error)
	}
)
