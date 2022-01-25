package service

import (
	"context"
	"encoding/json"
	"reflect"
)

type cache struct {
	rpc   RPC
	cache map[string]interface{}
}

// NewCache returns a cache wrapper around an RPC service.
func NewCache(rpc RPC) RPC {
	return &cache{
		rpc:   rpc,
		cache: make(map[string]interface{}),
	}
}

// Do an RPC request with cache enabled.
func (c *cache) Do(ctx context.Context, method string, out interface{}, args ...interface{}) error {
	return c.do(ctx, request{method: method, args: args}, out)
}

func (c *cache) do(ctx context.Context, r request, out interface{}) error {
	if v, ok := c.cache[r.Key()]; ok && out != nil {
		return c.write(v, out)
	}
	if err := c.rpc.Do(ctx, r.method, out, r.args...); err != nil {
		return err
	}

	c.cache[r.Key()] = out
	return nil
}

func (c *cache) write(dest, src interface{}) error {
	drv := reflect.ValueOf(dest)
	if drv.Kind() != reflect.Ptr || drv.IsNil() {
		return &json.InvalidUnmarshalError{Type: reflect.TypeOf(dest)}
	}

	for drv.Kind() == reflect.Ptr {
		drv = reflect.Indirect(drv)
	}

	srv := reflect.ValueOf(src)
	for srv.Kind() == reflect.Ptr {
		srv = reflect.Indirect(srv)
	}

	drv.Set(srv)

	return nil
}
