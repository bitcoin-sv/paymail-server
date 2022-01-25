package models

import (
	"fmt"
)

type blockVerbosity string

// Block verbosity levels.
const (
	VerbosityRawBlock                blockVerbosity = "RAW_BLOCK"
	VerbosityDecodeHeader            blockVerbosity = "DECODE_HEADER"
	VerbosityDecodeTransactions      blockVerbosity = "DECODE_TRANSACTIONS"
	VerbosityDecodeHeaderAndCoinbase blockVerbosity = "DECODE_HEADER_AND_COINBASE"
)

type merkleProofTargetType string

// Merkle proof target types.
const (
	MerkleProofTargetTypeHash       merkleProofTargetType = "hash"
	MerkleProofTargetTypeHeader     merkleProofTargetType = "header"
	MerkleProofTargetTypeMerkleRoot merkleProofTargetType = "merkleroot"
)

// Request model.
type Request struct {
	ID      string        `json:"id"`
	JSONRpc string        `json:"jsonRpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params,omitempty"`
}

// Response model.
type Response struct {
	Result interface{} `json:"result"`
	Error  *Error      `json:"error"`
}

// Error model.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

// OptsChainTxStats options.
type OptsChainTxStats struct {
	NumBlocks uint32
	BlockHash string
}

// Args convert struct into optional positional arguments.
func (o *OptsChainTxStats) Args() []interface{} {
	aa := []interface{}{o.NumBlocks}
	if o.BlockHash != "" {
		aa = append(aa, o.BlockHash)
	}
	return aa
}

// OptsMerkleProof options.
type OptsMerkleProof struct {
	FullTx     bool
	TargetType merkleProofTargetType
}

// Args convert struct into optional positional arguments.
func (o *OptsMerkleProof) Args() []interface{} {
	aa := []interface{}{o.FullTx}
	if o.TargetType != "" {
		aa = append(aa, o.TargetType)
	}

	return aa
}

// OptsLegacyMerkleProof options.
type OptsLegacyMerkleProof struct {
	BlockHash string
}

// Args convert struct into optional positional arguments.
func (o *OptsLegacyMerkleProof) Args() []interface{} {
	return []interface{}{o.BlockHash}
}

// OptsGenerate options.
type OptsGenerate struct {
	MaxTries uint32
}

// Args convert struct into optional positional arguments.
func (o *OptsGenerate) Args() []interface{} {
	return []interface{}{o.MaxTries}
}
