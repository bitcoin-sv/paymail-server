package bn

import (
	"context"
	"time"

	"github.com/libsv/go-bn/models"
)

// ControlClient interfaces interaction with the control sub commands on a bitcoin node.
type ControlClient interface {
	ActiveZMQNotifications(ctx context.Context) ([]*models.ZMQNotification, error)
	DumpParams(ctx context.Context) ([]string, error)
	Info(ctx context.Context) (*models.Info, error)
	MemoryInfo(ctx context.Context) (*models.MemoryInfo, error)
	Settings(ctx context.Context) (*models.Settings, error)
	Stop(ctx context.Context) error
	Uptime(ctx context.Context) (time.Duration, error)
}

// NewControlClient returns a client only capable of interfacing with the control sub commands on a bitcoin node.
func NewControlClient(oo ...BitcoinClientOptFunc) ControlClient {
	return NewNodeClient(oo...)
}

func (c *client) ActiveZMQNotifications(ctx context.Context) ([]*models.ZMQNotification, error) {
	var resp []*models.ZMQNotification
	return resp, c.rpc.Do(ctx, "activezmqnotifications", &resp)
}

func (c *client) DumpParams(ctx context.Context) ([]string, error) {
	var resp []string
	return resp, c.rpc.Do(ctx, "dumpparameters", &resp)
}

func (c *client) Info(ctx context.Context) (*models.Info, error) {
	var resp models.Info
	return &resp, c.rpc.Do(ctx, "getinfo", &resp)
}

func (c *client) MemoryInfo(ctx context.Context) (*models.MemoryInfo, error) {
	var resp models.MemoryInfo
	return &resp, c.rpc.Do(ctx, "getmemoryinfo", &resp)
}

func (c *client) Settings(ctx context.Context) (*models.Settings, error) {
	var resp models.Settings
	return &resp, c.rpc.Do(ctx, "getsettings", &resp)
}

func (c *client) Stop(ctx context.Context) error {
	return c.rpc.Do(ctx, "stop", nil)
}

func (c *client) Uptime(ctx context.Context) (time.Duration, error) {
	var resp time.Duration
	return resp, c.rpc.Do(ctx, "uptime", &resp)
}
