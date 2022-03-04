package data

import (
	"context"
	"net/http"

	"github.com/nch-bowstave/paymail/config"
	"github.com/nch-bowstave/paymail/models"
)

// DPP interfaces interactions with a dpp server.
type DPP interface {
	PaymentRequest(ctx context.Context, req models.DPPPayRequest) (*models.PaymentRequest, error)
	PaymentSend(ctx context.Context, args models.DPPPayRequest, req models.Payment) (*models.PaymentACK, error)
	Host() string
}

type dppClient struct {
	c    HTTPClient
	host *config.DPP
}

// NewDPP returns a new dpp interface.
func NewDPP(cfg *config.DPP, c HTTPClient) DPP {
	return &dppClient{
		c:    c,
		host: cfg,
	}
}

func (p *dppClient) Host() string {
	return p.host.ServerHost
}

// PaymentRequest performs a payment request http request to the specified url.
func (p *dppClient) PaymentRequest(ctx context.Context, args models.DPPPayRequest) (*models.PaymentRequest, error) {
	var resp models.PaymentRequest
	err := p.c.Do(ctx, http.MethodGet, args.PayToURL, http.StatusOK, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PaymentSend sends a payment http request to the specified url, with the provided payment packet.
func (p *dppClient) PaymentSend(ctx context.Context, args models.DPPPayRequest, req models.Payment) (*models.PaymentACK, error) {
	var resp models.PaymentACK
	err := p.c.Do(ctx, http.MethodPost, args.PayToURL, http.StatusAccepted, &req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
