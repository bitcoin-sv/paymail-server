package models

// MultiSig model.
type MultiSig struct {
	Address      string `json:"address"`
	RedeemScript string `json:"redeemScript"`
}

// ValidateAddress model.
type ValidateAddress struct {
	IsValid       bool   `json:"isvalid"`
	Address       string `json:"address"`
	LockingScript string `json:"scriptPubKey"`
	IsMine        bool   `json:"ismine"`
	IsWatchOnly   bool   `json:"iswatchonly"`
	IsScript      bool   `json:"isscript"`
	PublicKey     string `json:"pubkey"`
	IsCompressed  bool   `json:"iscompressed"`
	Account       string `json:"account"`
	Timestamp     uint64 `json:"timestamp"`
	HDKeyPath     string `json:"hdkeypath"`
	HDMasterKeyID string `json:"hdmasterkeyid"`
}
