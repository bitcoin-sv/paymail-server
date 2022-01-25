package service

import (
	"context"
	"fmt"
)

// RPC interface with an rpc server.
type RPC interface {
	Do(ctx context.Context, method string, out interface{}, args ...interface{}) error
}

type request struct {
	method string
	args   []interface{}
}

func (r request) Key() string {
	return fmt.Sprintf("%s|%s", r.method, r.args)
}
