package models

// ChainInfo model.
type ChainInfo struct {
	Chain                string  `json:"chain"`
	Blocks               uint32  `json:"blocks"`
	Headers              uint32  `json:"headers"`
	BestBlockHash        string  `json:"bestblockhash"`
	Difficulty           float64 `json:"difficulty"`
	MedianTime           uint32  `json:"mediantime"`
	VerificationProgress uint8   `json:"verificationprogress"`
	Chainwork            string  `json:"chainwork"`
	Pruned               bool    `json:"pruned"`
	SoftForks            []struct {
		ID      string `json:"id"`
		Version uint32 `json:"version"`
		Reject  struct {
			Status bool `json:"status"`
		} `json:"reject"`
	} `json:"softforks"`
}

// BlockStats model.
type BlockStats struct {
	AvgFee           float64 `json:"avgfee"`
	AvgFeeRate       float64 `json:"avgfeerate"`
	AvgTxSize        uint32  `json:"avgtxsize"`
	Blockhash        string  `json:"blockhash"`
	Height           uint32  `json:"height"`
	Ins              uint32  `json:"ins"`
	MaxFee           float64 `json:"maxfee"`
	MaxFeeRate       float64 `json:"maxfeerate"`
	MaxTxSize        uint32  `json:"maxtxsize"`
	MedianFee        float64 `json:"medianfee"`
	MedianFeeRate    float64 `json:"medianfeerate"`
	MedianTime       uint32  `json:"mediantime"`
	MedianTxSize     uint32  `json:"mediantxsize"`
	MinFee           float64 `json:"minfee"`
	MinFeeRate       float64 `json:"minfeerate"`
	MinTxSize        uint32  `json:"mintxsize"`
	Outs             uint32  `json:"outs"`
	Subsidy          float64 `json:"subsidy"`
	Time             uint32  `json:"time"`
	TotalOut         float64 `json:"total_out"`
	TotalSize        uint64  `json:"total_size"`
	TotalFee         float64 `json:"totalfee"`
	UtxoIncreate     uint32  `json:"utxo_increase"`
	UtxoSizeIncrease uint32  `json:"utxo_size_inc"`
}

// ChainTip model.
type ChainTip struct {
	Height    uint32 `json:"height"`
	Hash      string `json:"hash"`
	BranchLen uint32 `json:"branchLen"`
	Status    string `json:"status"`
}

// ChainTxStats model.
type ChainTxStats struct {
	Time             uint32  `json:"time"`
	TxCount          uint32  `json:"txcount"`
	WindowBlockCount uint32  `json:"window_block_count"`
	WindowTxCount    uint32  `json:"window_tx_count"`
	WindowInterval   uint32  `json:"window_interval"`
	TxRate           float32 `json:"txrate"`
}

// LegacyMerkleProof model.
type LegacyMerkleProof struct {
	Flags  uint32      `json:"flags"`
	Index  uint32      `json:"index"`
	TxOrID string      `json:"txOrId"`
	Target BlockHeader `json:"target"`
	Nodes  []string    `json:"nodes"`
}

// MempoolEntry model.
type MempoolEntry struct {
	Size        uint32   `json:"size"`
	Fee         float64  `json:"fee"`
	ModifiedFee float64  `json:"modifiedfee"`
	Time        uint32   `json:"time"`
	Height      uint32   `json:"height"`
	Depends     []string `json:"depends"`
}

// MempoolTxs model.
type MempoolTxs map[string]MempoolEntry

// JournalStatus model.
type JournalStatus struct {
	Ok     bool    `json:"ok"`
	Errors *string `json:"errors"`
}
