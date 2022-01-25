package service

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"golang.org/x/sync/singleflight"

	"github.com/libsv/go-bn/internal/config"
	"github.com/libsv/go-bn/models"
	"github.com/pkg/errors"
)

// Boiler RPC fields.
const (
	ID      = "go-bn"
	JSONRpc = "1.0"
)

var (
	// ErrRPCQuery error when rpc query fails.
	ErrRPCQuery = errors.New("failed to perform rpc query")
)

type rpc struct {
	c   *http.Client
	cfg *config.RPC
	g   singleflight.Group
}

// NewRPC returns a new RPC configured RPC client.
func NewRPC(cfg *config.RPC, c *http.Client) RPC {
	return &rpc{
		cfg: cfg,
		c:   c,
		g:   singleflight.Group{},
	}
}

// Do an RPC request.
func (h *rpc) Do(ctx context.Context, method string, out interface{}, args ...interface{}) error {
	return h.do(ctx, request{method: method, args: args}, out)
}

func (h *rpc) do(ctx context.Context, r request, out interface{}) error {
	data, err, _ := h.g.Do(r.Key(), func() (interface{}, error) {
		data, err := json.Marshal(&models.Request{
			ID:      ID,
			JSONRpc: JSONRpc,
			Method:  r.method,
			Params:  r.args,
		})
		if err != nil {
			return nil, err
		}

		req, err := http.NewRequestWithContext(
			ctx,
			http.MethodPost,
			h.cfg.Host,
			bytes.NewReader(data),
		)
		if err != nil {
			return nil, err
		}
		req.SetBasicAuth(h.cfg.Username, h.cfg.Password)
		req.Header.Add("Content-Type", "text/plain")

		resp, err := h.c.Do(req)
		if err != nil {
			return nil, err
		}
		defer func() {
			_ = resp.Body.Close()
		}()

		bb, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return bb, nil
	})
	if err != nil {
		return err
	}

	if v, ok := out.(interface {
		NodeJSON() interface{}
	}); ok {
		out = v.NodeJSON()
	}

	reply := models.Response{
		Result: out,
	}
	if err = json.NewDecoder(bytes.NewBuffer(data.([]byte))).Decode(&reply); err != nil {
		return err
	}

	if reply.Error != nil {
		return reply.Error
	}

	if v, ok := out.(interface {
		PostProcess() error
	}); ok {
		return v.PostProcess()
	}

	return nil
}
