package paymail

// Capabilities is a struct detailing the supported capability endpoints.
type Capabilities struct {
	PKI                string `json:"pki,omitempty"`                // Get public key information - Alternate: 0c4339ef99c2
	PaymentDestination string `json:"paymentDestination,omitempty"` // Resolve an address aka Payment Destination - Alternate: 759684b1a19a
	PublicProfile      string `json:"f12f968c92d6,omitempty"`       // Returns a public profile
	VerifyPublicKey    string `json:"a9f510c16bde,omitempty"`       // Verify a given pubkey
}

// Capability is the top level struct for the paymail capability. It contains a bsv alias and a list of capabilities.
type Capability struct {
	BSVAlias     string `json:"bsvalias"`
	Capabilities `json:"capabilities"`
}
