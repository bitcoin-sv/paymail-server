package models

import (
	"time"

	"github.com/libsv/go-bc/spv"
	"github.com/libsv/go-bt/v2"
	"gopkg.in/guregu/null.v3"

	"github.com/libsv/go-dpp"
)

// PayDPaymentRequest is used to send a payment to PayD for valdiation and storage.
type PayDPaymentRequest struct {
	SPVEnvelope    *spv.Envelope                `json:"spvEnvelope"`
	RawTX          *string                      `json:"rawTx"`
	ProofCallbacks map[string]dpp.ProofCallback `json:"proofCallbacks"`
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

// InvoiceArgs contains argument/s to return a single invoice.
type InvoiceArgs struct {
	InvoiceID string `param:"invoiceID" db:"invoice_id"`
}

// User information on wallet users.
type User struct {
	ID           uint64                 `json:"id" db:"user_id"`
	Name         string                 `json:"name" db:"name"`
	Email        string                 `json:"email" db:"email"`
	Avatar       string                 `json:"avatar" db:"avatar_url"`
	Address      string                 `json:"address" db:"address"`
	PhoneNumber  string                 `json:"phoneNumber" db:"phone_number"`
	ExtendedData map[string]interface{} `json:"extendedData"`
}

// UserDetails information on wallet user without an id.
type UserDetails struct {
	Name        string `json:"name" db:"name"`
	Email       string `json:"email" db:"email"`
	Avatar      string `json:"avatar" db:"avatar_url"`
	Address     string `json:"address" db:"address"`
	PhoneNumber string `json:"phoneNumber" db:"phone_number"`
}

// ProofCallback contains information relating to a merkleproof callback.
type ProofCallback struct {
	// Token to use for authentication when sending the proof to the destination. Optional.
	Token string
}

// PaymentCreate is submitted to validate and add a payment to the wallet.
type PaymentCreate struct {
	// MerchantData is copied from PaymentDetails.merchantData.
	// Payment hosts may use invoice numbers or any other data they require to match Payments to PaymentRequests.
	// Note that malicious clients may modify the merchantData, so should be authenticated
	// in some way (for example, signed with a payment host-only key).
	// Maximum length is 10000 characters.
	MerchantData User `json:"merchantData"`
	// RefundTo is a paymail to send a refund to should a refund be necessary.
	// Maximum length is 100 characters
	RefundTo null.String `json:"refundTo" swaggertype:"primitive,string" example:"me@paymail.com"`
	// Memo is a plain-text note from the customer to the payment host.
	Memo string `json:"memo" example:"for invoice 123456"`
	// SPVEnvelope which contains the details of previous transaction and Merkle proof of each input UTXO.
	// Should be available if SPVRequired is set to true in the paymentRequest.
	// See https://tsc.bitcoinassociation.net/standards/spv-envelope/
	SPVEnvelope *spv.Envelope `json:"spvEnvelope"`
	// ProofCallbacks are optional and can be supplied when the sender wants to receive
	// a merkleproof for the transaction they are submitting as part of the SPV Envelope.
	//
	// This is especially useful if they are receiving change and means when they use it
	// as an input, they can provide the merkle proof.
	ProofCallbacks map[string]ProofCallback `json:"proofCallbacks"`
}

// InvoiceCreate is used to create a new invoice.
type InvoiceCreate struct {
	InvoiceID string `json:"-" db:"invoice_id"`
	// Satoshis is the total amount this invoice is to pay.
	Satoshis uint64 `json:"satoshis" db:"satoshis"`
	// Reference is an identifier that can be used to link the
	// payd invoice with an external system.
	// MaxLength is 32 characters.
	Reference null.String `json:"reference" db:"payment_reference" swaggertype:"primitive,string"`
	// Description is an optional text field that can have some further info
	// like 'invoice for oranges'.
	// MaxLength is 1024 characters.
	Description null.String `json:"description" db:"description" swaggertype:"primitive,string"`
	// CreatedAt is the timestamp when the invoice was created.
	CreatedAt time.Time `json:"-" db:"created_at"`
	// ExpiresAt is an optional param that can be passed to set an expiration
	// date on an invoice, after which, payments will not be accepted.
	ExpiresAt null.Time `json:"expiresAt" db:"expires_at"`
	// SPVRequired if true will mean this invoice requires a valid spvenvelope otherwise a rawTX will suffice.
	SPVRequired bool `json:"-" db:"spv_required"`
	// UserID should optionally address a particular user in the payd database which this invoice ought to be associated with.
	UserID uint64 `json:"user_id,omitempty" db:"user_id,omitempty"`
}

// InvoiceState enforces invoice states.
type InvoiceState string

// contains states that an invocie can have.
const (
	StateInvoicePending  InvoiceState = "pending"
	StateInvoicePaid     InvoiceState = "paid"
	StateInvoiceRefunded InvoiceState = "refunded"
	StateInvoiceDeleted  InvoiceState = "deleted"
)

func (i InvoiceState) String() string {
	return string(i)
}

// Invoice identifies a single payment request from this payd wallet,
// it states the amount, id and optional refund address. This indicate
// we are requesting n satoshis in payment.
type Invoice struct {
	// ID is a unique identifier for an invoice and can be used
	// to lookup a single invoice.
	ID string `json:"id" db:"invoice_id"`
	// Reference is an identifier that can be used to link the
	// PayD invoice with an external system.
	Reference null.String `json:"reference" db:"payment_reference"`
	// Description is an optional text field that can have some further info
	// like 'invoice for oranges'.
	Description null.String `json:"description" db:"description"`
	// Satoshis is the total amount this invoice is to pay.
	Satoshis uint64 `json:"satoshis" db:"satoshis"`
	// ExpiresAt is an optional param that can be passed to set an expiration
	// date on an invoice, after which, payments will not be accepted.
	ExpiresAt null.Time `json:"expiresAt" db:"expires_at"`
	// PaymentReceivedAt will be set when this invoice has been paid and
	// states when the payment was received in UTC time.
	PaymentReceivedAt null.Time `json:"paymentReceivedAt" db:"payment_received_at"`
	// RefundTo is an optional paymail address that can be used to refund the
	// customer if required.
	RefundTo null.String `json:"refundTo" db:"refund_to"`
	// RefundedAt if this payment has been refunded, this date will be set
	// to the UTC time of the refund.
	RefundedAt null.Time `json:"refundedAt" db:"refunded_at"`
	// State is the current status of the invoice.
	State InvoiceState `json:"state" db:"state" enums:"pending,paid,refunded,deleted"`
	// SPVRequired if true will mean this invoice requires a valid spvenvelope otherwise a rawTX will suffice.
	SPVRequired bool `json:"-" db:"spv_required"`
	MetaData
}

// MetaData contains common meta info for objects.
type MetaData struct {
	// CreatedAt is the UTC time the object was created.
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	// UpdatedAt is the UTC time the object was updated.
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
	// DeletedAt is the date the object was removed.
	DeletedAt null.Time `json:"deletedAt,omitempty" db:"deleted_at"`
}
