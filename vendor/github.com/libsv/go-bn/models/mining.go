package models

// MiningCandidate model.
type MiningCandidate struct {
	ID                  string   `json:"id"`
	PrevHash            string   `json:"prevhash"`
	Coinbase            string   `json:"coinbase,omitempty"`
	CoinbaseValue       uint64   `json:"coinbaseValue"`
	Version             uint64   `json:"version"`
	NBits               string   `json:"nBits"`
	Time                uint64   `json:"time"`
	Height              uint64   `json:"height"`
	NumTx               uint64   `json:"num_tx"`
	SizeWithoutCoinbase uint64   `json:"sizeWithoutCoinbase"`
	MerkleProofs        []string `json:"merkleProofs"`
}

// MiningInfo model.
type MiningInfo struct {
	Blocks           uint64  `json:"blocks"`
	CurrentBlockSize uint64  `json:"currentblocksize"`
	CurrentBlockTx   uint64  `json:"currentblocktx"`
	Difficulty       float64 `json:"difficulty"`
	Errors           string  `json:"errors"`
	NetworkHashPS    float64 `json:"networkhashps"`
	PooledTx         uint64  `json:"pooledtx"`
	Chain            string  `json:"chain"`
}

// MiningSolution model.
type MiningSolution struct {
	ID       string `json:"id"`
	Nonce    uint64 `json:"nonce"`
	Coinbase string `json:"coinbase,omitempty"`
	Time     uint64 `json:"time,omitempty"`
	Version  uint64 `json:"version,omitempty"`
}

// OptsMiningCandidate options.
type OptsMiningCandidate struct {
	IncludeCoinbase bool
}

// Args convert struct into optional positional arguments.
func (o *OptsMiningCandidate) Args() []interface{} {
	return []interface{}{o.IncludeCoinbase}
}

// OptsNetworkHashPS options.
type OptsNetworkHashPS struct {
	NumBlocks uint64
	Height    uint64
}

// Args convert struct into optional positional arguments.
func (o *OptsNetworkHashPS) Args() []interface{} {
	aa := []interface{}{o.NumBlocks}
	if o.Height != 0 {
		aa = append(aa, o.Height)
	}

	return aa
}

// OptsSubmitBlock options.
type OptsSubmitBlock struct {
	WorkID string `json:"workid"`
}

// Args convert struct into optional positional arguments.
func (o *OptsSubmitBlock) Args() []interface{} {
	return []interface{}{o}
}
