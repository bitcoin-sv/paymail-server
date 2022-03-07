package models

import (
	"time"

	"github.com/libsv/go-bc/spv"
	"github.com/libsv/go-bt/v2"
)

type DPPPayRequest struct {
	PayToURL string `json:"payToURL"`
}

// Output message used in BIP270.
// See https://github.com/moneybutton/bips/blob/master/bip-0270.mediawiki#output
type Output struct {
	// Amount is the number of satoshis to be paid.
	Amount uint64 `json:"amount" example:"100000"`
	// Script is a locking script where payment should be sent, formatted as a hexadecimal string.
	Script string `json:"script" example:"76a91455b61be43392125d127f1780fb038437cd67ef9c88ac"`
	// Description, an optional description such as "tip" or "sales tax". Maximum length is 100 chars.
	Description string `json:"description" example:"paymentReference 123456"`
}

// PaymentDestinations contains the supported destinations
// by this DPP server.
type PaymentDestinations struct {
	Outputs []Output `json:"outputs"`
}

// Destinations message containing outputs and their fees.
type Destinations struct {
	SPVRequired bool         `json:"spvRequired"`
	Network     string       `json:"network"`
	Outputs     []Output     `json:"outputs"`
	Fees        *bt.FeeQuote `json:"fees"`
	CreatedAt   time.Time    `json:"createdAt"`
	ExpiresAt   time.Time    `json:"expiresAt"`
}

// PaymentRequest message used in BIP270.
// See https://github.com/moneybutton/bips/blob/master/bip-0270.mediawiki#paymentrequest
type PaymentRequest struct {
	// Network  Always set to "bitcoin" (but seems to be set to 'bitcoin-sv'
	// outside bip270 spec, see https://handcash.github.io/handcash-merchant-integration/#/merchant-payments)
	// {enum: bitcoin, bitcoin-sv, test}
	// Required.
	Network string `json:"network" example:"mainnet" enums:"mainnet,testnet,stn,regtest"`
	// SPVRequired if true will expect the sender to submit an SPVEnvelope in the payment request, otherwise
	// a rawTx will be required.
	SPVRequired bool `json:"spvRequired" example:"true"`
	// Destinations contains supported payment destinations by the merchant and dpp server, initial P2PKH outputs but can be extended.
	// Required.
	Destinations PaymentDestinations `json:"destinations"`
	// CreationTimestamp Unix timestamp (seconds since 1-Jan-1970 UTC) when the PaymentRequest was created.
	// Required.
	CreationTimestamp time.Time `json:"creationTimestamp" swaggertype:"primitive,string" example:"2019-10-12T07:20:50.52Z"`
	// ExpirationTimestamp Unix timestamp (UTC) after which the PaymentRequest should be considered invalid.
	// Optional.
	ExpirationTimestamp time.Time `json:"expirationTimestamp" swaggertype:"primitive,string" example:"2019-10-12T07:20:50.52Z"`
	// PaymentURL secure HTTPS location where a Payment message (see below) will be sent to obtain a PaymentACK.
	// Maximum length is 4000 characters
	PaymentURL string `json:"paymentUrl" example:"https://localhost:3443/api/v1/payment/123456"`
	// Memo Optional note that should be displayed to the customer, explaining what this PaymentRequest is for.
	// Maximum length is 50 characters.
	Memo string `json:"memo" example:"invoice number 123456"`
	// MerchantData contains arbitrary data that may be used by the payment host to identify the PaymentRequest.
	// May be omitted if the payment host does not need to associate Payments with PaymentRequest
	// or if they associate each PaymentRequest with a separate payment address.
	// Maximum length is 10000 characters.
	MerchantData *Merchant `json:"merchantData,omitempty"`
	// FeeRate defines the amount of fees a users wallet should add to the payment
	// when submitting their final payments.
	FeeRate *bt.FeeQuote `json:"fees"`
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
