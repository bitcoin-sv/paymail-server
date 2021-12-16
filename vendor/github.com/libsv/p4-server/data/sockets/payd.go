package sockets

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/libsv/go-bk/envelope"
	"github.com/pkg/errors"
	"github.com/theflyingcodr/sockets"

	"github.com/libsv/go-p4"
)

// Routes contain the unique keys for socket messages used in the payment protocol.
const (
	RoutePayment                = "payment"
	RoutePaymentACK             = "payment.ack"
	RouteProofCreate            = "proof.create"
	RoutePaymentRequestCreate   = "paymentrequest.create"
	RoutePaymentRequestResponse = "paymentrequest.response"
)

type payd struct {
	s sockets.ServerChannelBroadcaster
}

// NewPayd will setup and return a new payd socket data store.
func NewPayd(b sockets.ServerChannelBroadcaster) *payd {
	return &payd{s: b}
}

// ProofCreate will broadcast the proof to all currently listening clients on the socket channel.
func (p *payd) ProofCreate(ctx context.Context, args p4.ProofCreateArgs, req envelope.JSONEnvelope) error {
	msg := sockets.NewMessage("proof.create", "", args.PaymentReference)
	msg.AppID = "p4"
	msg.CorrelationID = args.TxID
	if err := msg.WithBody(req); err != nil {
		return err
	}
	msg.Headers.Add("x-tx-id", args.TxID)
	p.s.Broadcast(args.PaymentReference, msg)
	return nil
}

// PaymentRequest will send a socket request to a payd client for a payment request.
// It will wait on a response before returnign the payment request.
func (p *payd) PaymentRequest(ctx context.Context, args p4.PaymentRequestArgs) (*p4.PaymentRequest, error) {
	msg := sockets.NewMessage(RoutePaymentRequestCreate, "", args.PaymentID)
	msg.AppID = "p4"
	msg.CorrelationID = uuid.NewString()

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	resp, err := p.s.BroadcastAwait(ctx, args.PaymentID, msg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to broadcast message for payment request")
	}
	var pr *p4.PaymentRequest
	if err := resp.Bind(&pr); err != nil {
		return nil, errors.Wrap(err, "failed to bind payment request response")
	}
	return pr, nil
}

// PaymentCreate will send a request to payd to create and process the payment.
func (p *payd) PaymentCreate(ctx context.Context, args p4.PaymentCreateArgs, req p4.Payment) (*p4.PaymentACK, error) {
	msg := sockets.NewMessage(RoutePayment, "", args.PaymentID)
	msg.AppID = "p4"
	msg.CorrelationID = uuid.NewString()
	if err := msg.WithBody(req); err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	resp, err := p.s.BroadcastAwait(ctx, args.PaymentID, msg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send payment message for payment")
	}
	var pr *p4.PaymentACK
	if err := resp.Bind(&pr); err != nil {
		return nil, errors.Wrap(err, "failed to bind payment ack response")
	}
	return pr, nil
}
