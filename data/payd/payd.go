package payd

import (
	"context"
	"fmt"
	"net/http"

	"github.com/libsv/go-bk/envelope"
	"github.com/libsv/go-bt/v2/bscript"
	"github.com/pkg/errors"

	"github.com/libsv/dpp-proxy/data"
	"github.com/libsv/go-dpp"
	"github.com/nch-bowstave/paymail/config"
	"github.com/nch-bowstave/paymail/models"
)

// Known endpoints for the payd wallet implementing the payment protocol interface.
const (
	urlPayments      = "%s/api/v1/payments/%s"
	urlUser          = "%s/api/v1/users/%s"
	urlUserCreate    = "%s/api/v1/users"
	urlCreate        = "%s/api/v1/invoices"
	urlGet           = "%s/api/v1/invoices/%s"
	urlDestinations  = "%s/api/v1/destinations/%s"
	urlProofs        = "%s/api/v1/proofs/%s"
	protocolInsecure = "http"
	protocolSecure   = "https"
)

type Payd struct {
	client data.HTTPClient
	cfg    *config.PayD
}

// NewPayD will setup a new store that can interface with a payd wallet implementing
// the Payment Protocol Interface.
func NewPayD(cfg *config.PayD, client data.HTTPClient) *Payd {
	return &Payd{
		cfg:    cfg,
		client: client,
	}
}

// PaymentCreate will post a request to payd to validate and add the txos to the wallet.
//
// If invalid a non 204 status code is returned.
func (p *Payd) PaymentCreate(ctx context.Context, args dpp.PaymentCreateArgs, req dpp.Payment) (*dpp.PaymentACK, error) {
	paymentReq := models.PayDPaymentRequest{
		RawTX:          req.RawTX,
		SPVEnvelope:    req.SPVEnvelope,
		ProofCallbacks: req.ProofCallbacks,
	}
	if err := p.client.Do(ctx, http.MethodPost, fmt.Sprintf(urlPayments, p.baseURL(), args.PaymentID), http.StatusNoContent, paymentReq, nil); err != nil {
		return nil, err
	}
	return &dpp.PaymentACK{
		ID:   "",
		TxID: "",
		Memo: req.Memo,
		PeerChannel: &dpp.PeerChannelData{
			Host:      "",
			ChannelID: "",
			Token:     "",
		},
		Error: 0,
	}, nil
}

// User will return information regarding the owner of a payd wallet.
//
// In this example, the payd wallet has no auth, in proper implementations auth would
// be enabled and a cookie / oauth / bearer token etc would be passed down.
func (p *Payd) User(ctx context.Context, userID uint64) (*dpp.Merchant, error) {
	uid := fmt.Sprint(userID)
	var user *dpp.Merchant
	if err := p.client.Do(ctx, http.MethodGet, fmt.Sprintf(urlUser, p.baseURL(), uid), http.StatusOK, nil, &user); err != nil {
		return nil, errors.WithStack(err)
	}
	return user, nil
}

func (p *Payd) GetInvoiceByID(ctx context.Context, id string) (*models.Invoice, error) {
	var res models.Invoice
	if err := p.client.Do(ctx, http.MethodGet, fmt.Sprintf(urlGet, p.baseURL(), id), http.StatusOK, nil, &res); err != nil {
		return nil, errors.WithStack(err)
	}
	return &res, nil
}

func (p *Payd) CreateInvoice(ctx context.Context, req *models.InvoiceCreate) (*models.Invoice, error) {
	var res models.Invoice
	if err := p.client.Do(ctx, http.MethodPost, fmt.Sprintf(urlCreate, p.baseURL()), http.StatusCreated, &req, &res); err != nil {
		return nil, errors.WithStack(err)
	}
	return &res, nil
}

func (p *Payd) Destinations(ctx context.Context, args dpp.PaymentRequestArgs) (*dpp.Destinations, error) {
	var resp models.DestinationResponse
	if err := p.client.Do(ctx, http.MethodGet, fmt.Sprintf(urlDestinations, p.baseURL(), args.PaymentID), http.StatusOK, nil, &resp); err != nil {
		return nil, errors.WithStack(err)
	}
	dests := &dpp.Destinations{
		SPVRequired: resp.SPVRequired,
		Network:     resp.Network,
		Outputs:     make([]dpp.Output, 0),
		Fees:        resp.Fees,
		CreatedAt:   resp.CreatedAt,
		ExpiresAt:   resp.ExpiresAt,
	}
	for _, o := range resp.Outputs {
		s, _ := bscript.NewFromHexString(o.Script)
		dests.Outputs = append(dests.Outputs, dpp.Output{
			Amount:        o.Satoshis,
			LockingScript: s,
		})
	}

	return dests, nil
}

// ProofCreate will pass on the proof to a payd instance for storage.
func (p *Payd) ProofCreate(ctx context.Context, args dpp.ProofCreateArgs, req envelope.JSONEnvelope) error {
	return errors.WithStack(p.client.Do(ctx, http.MethodPost, fmt.Sprintf(urlProofs, p.baseURL(), args.TxID), http.StatusCreated, req, nil))
}

// baseURL will return http or https depending on if we're using TLS.
func (p *Payd) baseURL() string {
	if p.cfg.Secure {
		return fmt.Sprintf("%s://%s%s", protocolSecure, p.cfg.Host, p.cfg.Port)
	}
	return fmt.Sprintf("%s://%s%s", protocolInsecure, p.cfg.Host, p.cfg.Port)
}

func (p *Payd) CreateUser(ctx context.Context, req models.UserDetails) (*models.User, error) {
	var user *models.User
	if req.Name == "" || req.Email == "" {
		return nil, errors.New("must include name and email for user registration.")
	}
	if err := p.client.Do(ctx, http.MethodPost, fmt.Sprintf(urlUserCreate, p.baseURL()), http.StatusOK, &req, &user); err != nil {
		return nil, errors.WithStack(err)
	}
	return user, nil
}
