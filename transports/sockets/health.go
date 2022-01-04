package sockets

import (
	"context"

	"github.com/theflyingcodr/sockets"
	"github.com/theflyingcodr/sockets/server"
)

type health struct{}

// NewHealthHandler will setup and return a new instance of a health handler.
func NewHealthHandler() *health {
	return &health{}
}

// Register will register new handler/s with the socket server.
func (p *health) Register(s *server.SocketServer) {
	s.RegisterDirectHandler("health", p.ping)
}

func (p *health) ping(ctx context.Context, msg *sockets.Message) (*sockets.Message, error) {
	return msg, nil
}
