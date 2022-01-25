package bn

import (
	"context"

	"github.com/libsv/go-bc"
	"github.com/libsv/go-bn/models"
)

// MiningClient interfaces interaction with the mining sub commands on a bitcoin node.
type MiningClient interface {
	BlockTemplate(ctx context.Context, opts *models.BlockTemplateRequest) (*models.BlockTemplate, error)
	MiningCandidate(ctx context.Context, opts *models.OptsMiningCandidate) (*models.MiningCandidate, error)
	MiningInfo(ctx context.Context) (*models.MiningInfo, error)
	NetworkHashPS(ctx context.Context, opts *models.OptsNetworkHashPS) (uint64, error)
	PrioritiseTx(ctx context.Context, txID string, feeDelta int64) (bool, error)
	SubmitBlock(ctx context.Context, block *bc.Block, params *models.OptsSubmitBlock) (string, error)
	SubmitMiningSolution(ctx context.Context, solution *models.MiningSolution) (string, error)
	VerifyBlockCandidate(ctx context.Context, block *bc.Block, params *models.OptsSubmitBlock) (string, error)
}

// NewMiningClient returns a client only capable of interfacing with the mining sub commands on a bitcoin node.
func NewMiningClient(oo ...BitcoinClientOptFunc) MiningClient {
	return NewNodeClient(oo...)
}

func (c *client) BlockTemplate(ctx context.Context, opts *models.BlockTemplateRequest) (*models.BlockTemplate, error) {
	var resp models.BlockTemplate
	return &resp, c.rpc.Do(ctx, "getblocktemplate", &resp, c.argsFor(opts)...)
}

func (c *client) MiningCandidate(ctx context.Context,
	opts *models.OptsMiningCandidate) (*models.MiningCandidate, error) {
	var resp models.MiningCandidate
	return &resp, c.rpc.Do(ctx, "getminingcandidate", &resp, c.argsFor(opts)...)
}

func (c *client) MiningInfo(ctx context.Context) (*models.MiningInfo, error) {
	var resp models.MiningInfo
	return &resp, c.rpc.Do(ctx, "getmininginfo", &resp)
}

// TODO: verify is working with non-zero value
func (c *client) NetworkHashPS(ctx context.Context, opts *models.OptsNetworkHashPS) (uint64, error) {
	var resp int64
	return uint64(resp), c.rpc.Do(ctx, "getnetworkhashps", &resp, c.argsFor(opts)...)
}

func (c *client) PrioritiseTx(ctx context.Context, txID string, feeDelta int64) (bool, error) {
	var resp bool
	return resp, c.rpc.Do(ctx, "prioritisetx", &resp, txID, 0, feeDelta)
}

func (c *client) SubmitBlock(ctx context.Context, block *bc.Block, params *models.OptsSubmitBlock) (string, error) {
	var resp string
	return resp, c.rpc.Do(ctx, "submitblock", &resp, c.argsFor(params, block.String())...)
}

func (c *client) SubmitMiningSolution(ctx context.Context, solution *models.MiningSolution) (string, error) {
	var resp string
	return resp, c.rpc.Do(ctx, "submitminingsolution", &resp, solution)
}

func (c *client) VerifyBlockCandidate(ctx context.Context, block *bc.Block,
	params *models.OptsSubmitBlock) (string, error) {
	var resp string
	return resp, c.rpc.Do(ctx, "verifyblockcandidate", &resp, c.argsFor(params, block.String())...)
}
