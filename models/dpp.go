package models

import (
	"time"

	"github.com/libsv/go-bc/spv"
	"github.com/libsv/go-bt/v2"
)

type DPPPayRequest struct {
	PayToURL string `json:"payToURL"`
}

// DPPOutput an output matching what a dpp server expects.
type DPPOutput struct {
	Amount      uint64 `json:"amount"`
	Script      string `json:"script"`
	Description string `json:"description"`
}

// DPPDestination defines a dpp payment destination object.
type DPPDestination struct {
	Outputs []DPPOutput `json:"outputs"`
}

// PaymentRequestResponse
type PaymentRequest struct {
	Network             string         `json:"network"`
	Destinations        DPPDestination `json:"destinations"`
	CreationTimestamp   time.Time      `json:"creationTimestamp"`
	ExpirationTimestamp time.Time      `json:"expirationTimestamp"`
	PaymentURL          string         `json:"paymentURL"`
	Memo                string         `json:"memo"`
	MerchantData        User           `json:"merchantData"`
	Fee                 *bt.FeeQuote   `json:"fees"`
	AncestryRequired    bool           `json:"ancestryRequired" example:"true"`
}

// Merchant to be displayed to the user.
type Merchant struct {
	// AvatarURL displays a canonical url to a merchants avatar.
	AvatarURL string `json:"avatar" example:"http://url.com"`
	// Name is a human readable string identifying the merchant.
	Name string `json:"name" example:"merchant 1"`
	// Email can be sued to contact the merchant about this transaction.
	Email string `json:"email" example:"merchant@m.com"`
	// Address is the merchants store / head office address.
	Address string `json:"address" example:"1 the street, the town, B1 1AA"`
	// ExtendedData can be supplied if the merchant wishes to send some arbitrary data back to the wallet.
	ExtendedData map[string]interface{} `json:"extendedData,omitempty"`
}

// Payment is a Payment message used in BIP270.
// See https://github.com/moneybutton/bips/blob/master/bip-0270.mediawiki#payment
type Payment struct {
	// MerchantData is copied from PaymentDetails.merchantData.
	// Payment hosts may use invoice numbers or any other data they require to match Payments to PaymentRequests.
	// Note that malicious clients may modify the merchantData, so should be authenticated
	// in some way (for example, signed with a payment host-only key).
	// Maximum length is 10000 characters.
	MerchantData Merchant `json:"merchantData"`
	// RefundTo is a paymail to send a refund to should a refund be necessary.
	// Maximum length is 100 characters
	RefundTo *string `json:"refundTo"  swaggertype:"primitive,string" example:"me@paymail.com"`
	// Memo is a plain-text note from the customer to the payment host.
	Memo string `json:"memo" example:"for invoice 123456"`
	// SPVEnvelope which contains the details of previous transaction and Merkle proof of each input UTXO.
	// Should be available if SPVRequired is set to true in the paymentRequest.
	// See https://tsc.bitcoinassociation.net/standards/spv-envelope/
	SPVEnvelope *spv.Envelope `json:"spvEnvelope"`
	// RawTX should be sent if SPVRequired is set to false in the payment request.
	RawTX *string `json:"rawTx"`
	// ProofCallbacks are optional and can be supplied when the sender wants to receive
	// a merkleproof for the transaction they are submitting as part of the SPV Envelope.
	//
	// This is especially useful if they are receiving change and means when they use it
	// as an input, they can provide the merkle proof.
	ProofCallbacks map[string]ProofCallback `json:"proofCallbacks"`
}

// PaymentACK message used in BIP270.
// See https://github.com/moneybutton/bips/blob/master/bip-0270.mediawiki#paymentack
type PaymentACK struct {
	Payment PaymentCreate `json:"payment"`
	Memo    string        `json:"memo,omitempty"`
	// A number indicating why the transaction was not accepted. 0 or undefined indicates no error.
	// A 1 or any other positive integer indicates an error. The errors are left undefined for now;
	// it is recommended only to use “1” and to fill the memo with a textual explanation about why
	// the transaction was not accepted until further numbers are defined and standardised.
	Error int `json:"error,omitempty"`
}
