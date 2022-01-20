package data

import (
	"context"
	"net/http"

	"github.com/nch-bowstave/paymail/config"
	"github.com/nch-bowstave/paymail/models"
)

// P4 interfaces interactions with a p4 server.
type P4 interface {
	PaymentRequest(ctx context.Context, req models.P4PayRequest) (*models.PaymentRequest, error)
	PaymentSend(ctx context.Context, args models.P4PayRequest, req models.Payment) (*models.PaymentACK, error)
	Host() string
}

type p4Client struct {
	c    HTTPClient
	host *config.P4
}

// NewP4 returns a new p4 interface.
func NewP4(cfg *config.P4, c HTTPClient) P4 {
	return &p4Client{
		c:    c,
		host: cfg,
	}
}

func (p *p4Client) Host() string {
	return p.host.ServerHost
}

// PaymentRequest performs a payment request http request to the specified url.
func (p *p4Client) PaymentRequest(ctx context.Context, args models.P4PayRequest) (*models.PaymentRequest, error) {
	var resp models.PaymentRequest
	err := p.c.Do(ctx, http.MethodGet, args.PayToURL, http.StatusAccepted, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PaymentSend sends a payment http request to the specified url, with the provided payment packet.
func (p *p4Client) PaymentSend(ctx context.Context, args models.P4PayRequest, req models.Payment) (*models.PaymentACK, error) {
	var resp models.PaymentACK
	err := p.c.Do(ctx, http.MethodPost, args.PayToURL, http.StatusAccepted, &req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
