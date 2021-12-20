package service

import (
	"context"
	"encoding/json"

	"github.com/libsv/p4-server/log"
	"github.com/nch-bowstave/paymail/data"
)

type paymail struct {
	l log.Logger
}

// NewPaymail will create and return a new paymail service.
func NewPaymail(l log.Logger) *paymail {
	return &paymail{
		l: l,
	}
}

// Paymail contains the handlers for paymail service endpoints.
type Paymail interface {
	Capabilities(ctx context.Context) (*CapabilitiesDocument, error)
	// PKI(ctx context.Context, args Handle) (*Pki, error)
	// P2pDest(ctx context.Context, args ) (*PaymentDestination, error)
	// P2pRawTx(ctx context.Context, handle Handle, args SenderRequest) (*PaymentOutput, error)
}

// Capabilties will create the response from a static document and return it.
func (svc *paymail) Capabilities(ctx context.Context) (*CapabilitiesDocument, error) {
	data, err := data.CapabilitiesData.LoadStaticDocument()
	if err != nil {
		svc.l.Error(err, "capabilities.json document not found")
		return nil, err
	}
	caps := &CapabilitiesDocument{}
	err = json.Unmarshal(data, caps)
	if err != nil {
		return nil, err
	}
	return caps, err
}
