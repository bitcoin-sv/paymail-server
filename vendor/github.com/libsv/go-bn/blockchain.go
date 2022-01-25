package bn

import (
	"context"

	"github.com/libsv/go-bc"
	"github.com/libsv/go-bn/models"
	"github.com/libsv/go-bt/v2"
)

// BlockChainClient interfaces interaction with the blockchain sub commands on a bitcoin node.
type BlockChainClient interface {
	BestBlockHash(ctx context.Context) (string, error)
	BlockHex(ctx context.Context, hash string) (string, error)
	BlockHexByHeight(ctx context.Context, height int) (string, error)
	BlockDecodeHeader(ctx context.Context, hash string) (*models.BlockDecodeHeader, error)
	BlockDecodeHeaderByHeight(ctx context.Context, height int) (*models.BlockDecodeHeader, error)
	Block(ctx context.Context, hash string) (*models.Block, error)
	BlockByHeight(ctx context.Context, height int) (*models.Block, error)
	ChainInfo(ctx context.Context) (*models.ChainInfo, error)
	BlockCount(ctx context.Context) (uint32, error)
	BlockHash(ctx context.Context, height int) (string, error)
	BlockHeader(ctx context.Context, hash string) (*models.BlockHeader, error)
	BlockHeaderHex(ctx context.Context, hash string) (string, error)
	BlockStats(ctx context.Context, hash string, fields ...string) (*models.BlockStats, error)
	BlockStatsByHeight(ctx context.Context, height int, fields ...string) (*models.BlockStats, error)
	ChainTips(ctx context.Context) ([]*models.ChainTip, error)
	ChainTxStats(ctx context.Context, opts *models.OptsChainTxStats) (*models.ChainTxStats, error)
	Difficulty(ctx context.Context) (float64, error)
	MerkleProof(ctx context.Context, blockHash, txID string, opts *models.OptsMerkleProof) (*bc.MerkleProof, error)
	LegacyMerkleProof(ctx context.Context, txID string,
		opts *models.OptsLegacyMerkleProof) (*models.LegacyMerkleProof, error)
	RawMempool(ctx context.Context) (models.MempoolTxs, error)
	RawMempoolIDs(ctx context.Context) ([]string, error)
	RawNonFinalMempool(ctx context.Context) ([]string, error)
	MempoolEntry(ctx context.Context, txID string) (*models.MempoolEntry, error)
	MempoolAncestors(ctx context.Context, txID string) (models.MempoolTxs, error)
	MempoolAncestorIDs(ctx context.Context, txID string) ([]string, error)
	MempoolDescendants(ctx context.Context, txID string) (models.MempoolTxs, error)
	MempoolDescendantIDs(ctx context.Context, txID string) ([]string, error)
	Output(ctx context.Context, txID string, n int, opts *models.OptsOutput) (*models.Output, error)
	OutputSetInfo(ctx context.Context) (*models.OutputSetInfo, error)
	PreciousBlock(ctx context.Context, blockHash string) error
	PruneChain(ctx context.Context, height int) (uint32, error)
	CheckJournal(ctx context.Context) (*models.JournalStatus, error)
	RebuildJournal(ctx context.Context) error
	VerifyChain(ctx context.Context) (bool, error)
	Generate(ctx context.Context, n int, opts *models.OptsGenerate) ([]string, error)
	GenerateToAddress(ctx context.Context, n int, addr string, opts *models.OptsGenerate) ([]string, error)
}

// NewBlockChainClient returns a client only capable of interfacing with the blockchain sub commands on a bitcoin node.
func NewBlockChainClient(oo ...BitcoinClientOptFunc) BlockChainClient {
	return NewNodeClient(oo...)
}

func (c *client) BestBlockHash(ctx context.Context) (string, error) {
	var resp string
	return resp, c.rpc.Do(ctx, "getbestblockhash", &resp)
}

func (c *client) BlockHex(ctx context.Context, hash string) (string, error) {
	var resp string
	return resp, c.rpc.Do(ctx, "getblock", &resp, hash, models.VerbosityRawBlock)
}

func (c *client) BlockHexByHeight(ctx context.Context, height int) (string, error) {
	var resp string
	return resp, c.rpc.Do(ctx, "getblockbyheight", &resp, height, models.VerbosityRawBlock)
}

func (c *client) BlockDecodeHeader(ctx context.Context, hash string) (*models.BlockDecodeHeader, error) {
	var resp models.BlockDecodeHeader
	return &resp, c.rpc.Do(ctx, "getblock", &resp, hash, models.VerbosityDecodeHeader)
}

func (c *client) BlockDecodeHeaderByHeight(ctx context.Context, height int) (*models.BlockDecodeHeader, error) {
	var resp models.BlockDecodeHeader
	return &resp, c.rpc.Do(ctx, "getblockbyheight", &resp, height, models.VerbosityDecodeHeader)
}

func (c *client) Block(ctx context.Context, hash string) (*models.Block, error) {
	var resp models.Block
	return &resp, c.rpc.Do(ctx, "getblock", &resp, hash, models.VerbosityDecodeTransactions)
}

func (c *client) BlockByHeight(ctx context.Context, height int) (*models.Block, error) {
	var resp models.Block
	return &resp, c.rpc.Do(ctx, "getblockbyheight", &resp, height, models.VerbosityDecodeTransactions)
}

func (c *client) ChainInfo(ctx context.Context) (*models.ChainInfo, error) {
	var resp models.ChainInfo
	return &resp, c.rpc.Do(ctx, "getblockchaininfo", &resp)
}

func (c *client) BlockCount(ctx context.Context) (uint32, error) {
	var resp uint32
	return resp, c.rpc.Do(ctx, "getblockcount", &resp)
}

func (c *client) BlockHash(ctx context.Context, height int) (string, error) {
	var resp string
	return resp, c.rpc.Do(ctx, "getblockhash", &resp, height)
}

func (c *client) BlockHeader(ctx context.Context, hash string) (*models.BlockHeader, error) {
	resp := models.BlockHeader{BlockHeader: &bc.BlockHeader{}}
	return &resp, c.rpc.Do(ctx, "getblockheader", &resp, hash, true)
}

func (c *client) BlockHeaderHex(ctx context.Context, hash string) (string, error) {
	var resp string
	return resp, c.rpc.Do(ctx, "getblockheader", &resp, hash, false)
}

func (c *client) BlockStats(ctx context.Context, hash string, fields ...string) (*models.BlockStats, error) {
	var resp models.BlockStats
	return &resp, c.rpc.Do(ctx, "getblockstats", &resp, hash, fields)
}

func (c *client) BlockStatsByHeight(ctx context.Context, height int, fields ...string) (*models.BlockStats, error) {
	var resp models.BlockStats
	return &resp, c.rpc.Do(ctx, "getblockstatsbyheight", &resp, height, fields)
}

func (c *client) ChainTips(ctx context.Context) ([]*models.ChainTip, error) {
	var resp []*models.ChainTip
	return resp, c.rpc.Do(ctx, "getchaintips", &resp)
}

func (c *client) ChainTxStats(ctx context.Context, opts *models.OptsChainTxStats) (*models.ChainTxStats, error) {
	var resp models.ChainTxStats
	return &resp, c.rpc.Do(ctx, "getchaintxstats", &resp, c.argsFor(opts)...)
}

func (c *client) Difficulty(ctx context.Context) (float64, error) {
	var resp float64
	return resp, c.rpc.Do(ctx, "getdifficulty", &resp)
}

func (c *client) MempoolEntry(ctx context.Context, txID string) (*models.MempoolEntry, error) {
	var resp models.MempoolEntry
	return &resp, c.rpc.Do(ctx, "getmempoolentry", &resp, txID)
}

func (c *client) RawMempool(ctx context.Context) (models.MempoolTxs, error) {
	var resp models.MempoolTxs
	return resp, c.rpc.Do(ctx, "getrawmempool", &resp, true)
}

func (c *client) RawMempoolIDs(ctx context.Context) ([]string, error) {
	var resp []string
	return resp, c.rpc.Do(ctx, "getrawmempool", &resp, false)
}

func (c *client) RawNonFinalMempool(ctx context.Context) ([]string, error) {
	var resp []string
	return resp, c.rpc.Do(ctx, "getrawnonfinalmempool", &resp)
}

func (c *client) MempoolAncestors(ctx context.Context, txID string) (models.MempoolTxs, error) {
	var resp models.MempoolTxs
	return resp, c.rpc.Do(ctx, "getmempoolancestors", &resp, txID, true)
}

func (c *client) MempoolAncestorIDs(ctx context.Context, txID string) ([]string, error) {
	var resp []string
	return resp, c.rpc.Do(ctx, "getmempoolancestors", &resp, txID, false)
}

func (c *client) MempoolDescendants(ctx context.Context, txID string) (models.MempoolTxs, error) {
	var resp models.MempoolTxs
	return resp, c.rpc.Do(ctx, "getmempooldescendants", &resp, txID, true)
}

func (c *client) MempoolDescendantIDs(ctx context.Context, txID string) ([]string, error) {
	var resp []string
	return resp, c.rpc.Do(ctx, "getmempooldescendants", &resp, txID, false)
}

func (c *client) MerkleProof(ctx context.Context, blockHash, txID string,
	opts *models.OptsMerkleProof) (*bc.MerkleProof, error) {
	var resp bc.MerkleProof
	return &resp, c.rpc.Do(ctx, "getmerkleproof2", &resp, c.argsFor(opts, blockHash, txID)...)
}

func (c *client) LegacyMerkleProof(ctx context.Context, txID string,
	opts *models.OptsLegacyMerkleProof) (*models.LegacyMerkleProof, error) {
	var resp models.LegacyMerkleProof
	return &resp, c.rpc.Do(ctx, "getmerkleproof", &resp, c.argsFor(opts, txID)...)
}

func (c *client) Output(ctx context.Context, txID string, n int, opts *models.OptsOutput) (*models.Output, error) {
	resp := models.Output{Output: &bt.Output{}}
	return &resp, c.rpc.Do(ctx, "gettxout", &resp, c.argsFor(opts, txID, n)...)
}

func (c *client) OutputSetInfo(ctx context.Context) (*models.OutputSetInfo, error) {
	var resp models.OutputSetInfo
	return &resp, c.rpc.Do(ctx, "gettxoutsetinfo", &resp)
}

func (c *client) PreciousBlock(ctx context.Context, blockHash string) error {
	return c.rpc.Do(ctx, "preciousblock", nil, blockHash)
}

func (c *client) PruneChain(ctx context.Context, height int) (uint32, error) {
	var resp uint32
	return resp, c.rpc.Do(ctx, "pruneblockchain", &resp, height)
}

func (c *client) CheckJournal(ctx context.Context) (*models.JournalStatus, error) {
	var resp models.JournalStatus
	return &resp, c.rpc.Do(ctx, "checkjournal", &resp)
}

func (c *client) RebuildJournal(ctx context.Context) error {
	return c.rpc.Do(ctx, "rebuildjournal", nil)
}

func (c *client) VerifyChain(ctx context.Context) (bool, error) {
	var resp bool
	return resp, c.rpc.Do(ctx, "verifychain", &resp)
}

func (c *client) Generate(ctx context.Context, n int, opts *models.OptsGenerate) ([]string, error) {
	var resp []string
	return resp, c.rpc.Do(ctx, "generate", &resp, c.argsFor(opts, n)...)
}

func (c *client) GenerateToAddress(ctx context.Context, n int, addr string,
	opts *models.OptsGenerate) ([]string, error) {
	var resp []string
	return resp, c.rpc.Do(ctx, "generatetoaddress", &resp, c.argsFor(opts, n, addr)...)
}
