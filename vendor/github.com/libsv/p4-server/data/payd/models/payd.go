package models

import (
	"time"

	"github.com/libsv/go-bc/spv"
	"github.com/libsv/go-bt/v2"

	"github.com/libsv/go-p4"
)

// PayDPaymentRequest is used to send a payment to PayD for valdiation and storage.
type PayDPaymentRequest struct {
	SPVEnvelope    *spv.Envelope               `json:"spvEnvelope"`
	RawTX          *string                     `json:"rawTx"`
	ProofCallbacks map[string]p4.ProofCallback `json:"proofCallbacks"`
}

// Destination is a payment output with locking script.
type Destination struct {
	Script   string `json:"script"`
	Satoshis uint64 `json:"satoshis"`
}

// DestinationResponse is the response for the destinations api.
type DestinationResponse struct {
	SPVRequired bool          `json:"spvRequired"`
	Network     string        `json:"network"`
	Outputs     []Destination `json:"outputs"`
	Fees        *bt.FeeQuote  `json:"fees"`
	CreatedAt   time.Time     `json:"createdAt"`
	ExpiresAt   time.Time     `json:"expiresAt"`
}
