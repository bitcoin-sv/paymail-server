package models

import (
	"encoding/hex"
	"encoding/json"

	"github.com/libsv/go-bc"
	"github.com/libsv/go-bt/v2"
)

// BlockDecodeHeader model.
type BlockDecodeHeader struct {
	Txs []string `json:"tx"`
	BlockHeader
}

// Block model.
type Block struct {
	Txs bt.Txs `json:"tx"`
	BlockHeader
}

// UnmarshalJSON unmarshal response.
func (b *Block) UnmarshalJSON(bb []byte) error {
	bj := struct {
		Txs json.RawMessage `json:"tx"`
		BlockHeader
	}{}
	if err := json.Unmarshal(bb, &bj); err != nil {
		return err
	}

	var txs bt.Txs
	if err := json.Unmarshal(bj.Txs, txs.NodeJSON()); err != nil {
		return err
	}

	b.Txs = txs
	b.BlockHeader = bj.BlockHeader
	return nil
}

// BlockHeader model.
type BlockHeader struct {
	*bc.BlockHeader
	Hash          string `json:"hash"`
	Confirmations uint64 `json:"confirmations"`
	Height        uint64 `json:"height"`
	//Version           uint64  `json:"version"`
	VersionHex string `json:"versionHex"`
	NumTx      uint64 `json:"num_tx"`
	//Time              uint64  `json:"time"`
	MedianTime uint64 `json:"mediantime"`
	//Nonce             uint64  `json:"nonce"`
	//Bits       string  `json:"bits"`
	Difficulty float64 `json:"difficulty"`
	Chainwork  string  `json:"chainwork"`
	//PreviousBlockHash string  `json:"previousblockhash"`
	NextBlockHash string `json:"nextblockhash"`
}

// UnmarshalJSON unmarshal response.
func (b *BlockHeader) UnmarshalJSON(bb []byte) error {
	bh := struct {
		Hash              string  `json:"hash"`
		Confirmations     uint64  `json:"confirmations"`
		Height            uint64  `json:"height"`
		VersionHex        string  `json:"versionHex"`
		NumTx             uint64  `json:"num_tx"`
		MerkleRoot        string  `json:"merkleroot"`
		MedianTime        uint64  `json:"mediantime"`
		Difficulty        float64 `json:"difficulty"`
		Chainwork         string  `json:"chainwork"`
		NextBlockHash     string  `json:"nextblockhash"`
		PreviousBlockHash string  `json:"previousblockhash"`
	}{}

	if err := json.Unmarshal(bb, &bh); err != nil {
		return err
	}

	var blockHeader bc.BlockHeader
	err := json.Unmarshal(bb, &blockHeader)
	if err != nil {
		return err
	}

	blockHeader.HashMerkleRoot, err = hex.DecodeString(bh.MerkleRoot)
	if err != nil {
		return err
	}

	blockHeader.HashPrevBlock, err = hex.DecodeString(bh.PreviousBlockHash)
	if err != nil {
		return err
	}

	b.Hash = bh.Hash
	b.Confirmations = bh.Confirmations
	b.Height = bh.Height
	b.VersionHex = bh.VersionHex
	b.NumTx = bh.NumTx
	b.MedianTime = bh.MedianTime
	b.Difficulty = bh.Difficulty
	b.Chainwork = bh.Chainwork
	b.NextBlockHash = bh.NextBlockHash
	*b.BlockHeader = blockHeader
	return nil
}

// MarshalJSON marshal response.
func (b *BlockHeader) MarshalJSON() ([]byte, error) {
	cpy := *b
	cpy.BlockHeader = nil
	bh := struct {
		BlockHeader
		PreviousBlockHash string `json:"previousblockhash"`
		MerkleRoot        string `json:"merkleroot"`
		Version           uint32 `json:"version"`
		Nonce             uint32 `json:"nonce"`
	}{
		PreviousBlockHash: b.HashPrevBlockStr(),
		MerkleRoot:        b.HashMerkleRootStr(),
		Version:           b.Version,
		Nonce:             b.Nonce,
		BlockHeader:       cpy,
	}

	return json.Marshal(bh)
}

// BlockTemplate model.
type BlockTemplate struct {
	Capabilities      []string `json:"capabilities"`
	Version           uint64   `json:"version"`
	PreviousBlockHash string   `json:"previousblockhash"`
	Transactions      []string `json:"transactions"`
	CoinbaseAux       struct {
		Flags string `json:"flags"`
	} `json:"coinbaseaux"`
	CoinbaseValue uint64   `json:"coinbasevalue"`
	LongPollID    string   `json:"longpollid"`
	Target        string   `json:"target"`
	MinTime       uint64   `json:"mintime"`
	Mutable       []string `json:"mutable"`
	NonceRange    string   `json:"noncerange"`
	SizeLimit     uint64   `json:"sizelimit"`
	CurTime       uint64   `json:"curtime"`
	Bits          string   `json:"bits"`
	Height        uint64   `json:"height"`
}

// BlockTemplateRequest model.
type BlockTemplateRequest struct {
	Mode         string
	Capabilities []string
}

// Args convert struct into optional positional arguments.
func (r *BlockTemplateRequest) Args() []interface{} {
	return []interface{}{r}
}
