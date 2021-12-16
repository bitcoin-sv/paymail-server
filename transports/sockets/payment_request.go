package sockets

import (
	"github.com/theflyingcodr/sockets"

	"context"

	"github.com/theflyingcodr/sockets/server"
)

type paymentRequest struct {
}

// NewPaymentRequest will setup a new instance of a paymentRequest handler.
func NewPaymentRequest() *paymentRequest {
	return &paymentRequest{}
}

// Register will register new handler/s with the socket server.
func (p *paymentRequest) Register(s *server.SocketServer) {
	s.RegisterChannelHandler("paymentrequest.create", p.buildPaymentRequest)
	s.RegisterChannelHandler("paymentrequest.response", p.paymentRequestResponse)
}

// buildPaymentRequest will forward a paymentrequest message to all connected clients.
func (p *paymentRequest) buildPaymentRequest(ctx context.Context, msg *sockets.Message) (*sockets.Message, error) {
	return msg, nil
}

// buildPaymentRequest will forward a paymentrequest.response message to all connected clients.
func (p *paymentRequest) paymentRequestResponse(ctx context.Context, msg *sockets.Message) (*sockets.Message, error) {
	return msg, nil
}
