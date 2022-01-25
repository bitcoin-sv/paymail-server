package bn

import (
	"context"

	"github.com/libsv/go-bn/internal"
	"github.com/libsv/go-bn/models"
)

// NodeAdd enums.
const (
	NodeAddOneTry internal.NodeAddType = "onetry"
	NodeAddRemove internal.NodeAddType = "remove"
	NodeAddAdd    internal.NodeAddType = "add"
)

// BanAction enums.
const (
	BanActionAdd    internal.BanAction = "add"
	BanActionRemove internal.BanAction = "remove"
)

// NetworkClient interfaces interaction with the network sub commands on a bitcoin node.
type NetworkClient interface {
	Ping(ctx context.Context) error
	AddNode(ctx context.Context, node string, command internal.NodeAddType) error
	ClearBanned(ctx context.Context) error
	DisconnectNode(ctx context.Context, params models.ParamsDisconnectNode) error
	NodeInfo(ctx context.Context, opts *models.OptsNodeInfo) ([]*models.NodeInfo, error)
	ConnectionCount(ctx context.Context) (uint64, error)
	ExcessiveBlock(ctx context.Context) (*models.ExcessiveBlock, error)
	NetworkTotals(ctx context.Context) (*models.NetworkTotals, error)
	NetworkInfo(ctx context.Context) (*models.NetworkInfo, error)
	PeerInfo(ctx context.Context) ([]*models.PeerInfo, error)
	ListBanned(ctx context.Context) ([]*models.BannedSubnet, error)
	SetBan(ctx context.Context, subnet string, action internal.BanAction, opts *models.OptsSetBan) error
	SetBlockMaxSize(ctx context.Context, size uint64) (string, error)
	SetExcessiveBlock(ctx context.Context, size uint64) (string, error)
	SetNetworkActive(ctx context.Context, enabled bool) error
	SetTxPropagationFrequency(ctx context.Context, frequency uint64) error
}

// NewNetworkClient returns a client only capable of interfacing with the network sub commands on a bitcoin node.
func NewNetworkClient(oo ...BitcoinClientOptFunc) NetworkClient {
	return NewNodeClient(oo...)
}

func (c *client) Ping(ctx context.Context) error {
	return c.rpc.Do(ctx, "ping", nil)
}

func (c *client) AddNode(ctx context.Context, node string, command internal.NodeAddType) error {
	return c.rpc.Do(ctx, "addnode", nil, node, command)
}

func (c *client) ClearBanned(ctx context.Context) error {
	return c.rpc.Do(ctx, "clearbanned", nil)
}

func (c *client) DisconnectNode(ctx context.Context, params models.ParamsDisconnectNode) error {
	return c.rpc.Do(ctx, "disconnectnode", nil, params.Args()...)
}

func (c *client) NodeInfo(ctx context.Context, opts *models.OptsNodeInfo) ([]*models.NodeInfo, error) {
	var resp []*models.NodeInfo
	return resp, c.rpc.Do(ctx, "getaddednodeinfo", &resp, c.argsFor(opts)...)
}

func (c *client) ConnectionCount(ctx context.Context) (uint64, error) {
	var resp uint64
	return resp, c.rpc.Do(ctx, "getconnectioncount", &resp)
}

func (c *client) ExcessiveBlock(ctx context.Context) (*models.ExcessiveBlock, error) {
	var resp models.ExcessiveBlock
	return &resp, c.rpc.Do(ctx, "getexcessiveblock", &resp)
}

func (c *client) NetworkTotals(ctx context.Context) (*models.NetworkTotals, error) {
	var resp models.NetworkTotals
	return &resp, c.rpc.Do(ctx, "getnettotals", &resp)
}

func (c *client) NetworkInfo(ctx context.Context) (*models.NetworkInfo, error) {
	var resp models.NetworkInfo
	return &resp, c.rpc.Do(ctx, "getnetworkinfo", &resp)
}

func (c *client) PeerInfo(ctx context.Context) ([]*models.PeerInfo, error) {
	var resp []*models.PeerInfo
	return resp, c.rpc.Do(ctx, "getpeerinfo", &resp)
}

func (c *client) ListBanned(ctx context.Context) ([]*models.BannedSubnet, error) {
	var resp []*models.BannedSubnet
	return resp, c.rpc.Do(ctx, "listbanned", &resp)
}

func (c *client) SetBan(ctx context.Context, subnet string, action internal.BanAction, opts *models.OptsSetBan) error {
	return c.rpc.Do(ctx, "setban", nil, c.argsFor(opts, subnet, action)...)
}

func (c *client) SetBlockMaxSize(ctx context.Context, size uint64) (string, error) {
	var resp string
	return resp, c.rpc.Do(ctx, "setblockmaxsize", &resp, size)
}

func (c *client) SetExcessiveBlock(ctx context.Context, size uint64) (string, error) {
	var resp string
	return resp, c.rpc.Do(ctx, "setexcessiveblock", &resp, size)
}

func (c *client) SetNetworkActive(ctx context.Context, enabled bool) error {
	return c.rpc.Do(ctx, "setnetworkactive", nil, enabled)
}

// TODO: work out how to use this
func (c *client) SetTxPropagationFrequency(ctx context.Context, frequency uint64) error {
	panic("not implemented")
}
