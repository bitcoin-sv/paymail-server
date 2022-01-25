package bn

import (
	"net/http"
	"reflect"
	"time"

	"github.com/libsv/go-bn/internal/config"
	"github.com/libsv/go-bn/internal/service"
)

// NodeClient interfaces interacting with all commands on a bitcoin node.
type NodeClient interface {
	BlockChainClient
	ControlClient
	MiningClient
	NetworkClient
	TransactionClient
	UtilClient
	WalletClient
}

type positionalOptionalArgs interface {
	Args() []interface{}
}

type client struct {
	rpc       service.RPC
	isMainnet bool
}

// NewNodeClient returns a node client, built from the provided option funcs.
// This client is used for interfacing with the bitcoin node across all subcategories.
func NewNodeClient(oo ...BitcoinClientOptFunc) NodeClient {
	opts := &clientOpts{
		timeout:  30 * time.Second,
		host:     "http://localhost:8332",
		username: "bitcoin",
		password: "bitcoin",
	}
	for _, o := range oo {
		o(opts)
	}

	if opts.rpc != nil {
		return &client{
			rpc:       opts.rpc,
			isMainnet: opts.isMainnet,
		}
	}

	rpc := service.NewRPC(&config.RPC{
		Username: opts.username,
		Password: opts.password,
		Host:     opts.host,
	}, &http.Client{Timeout: opts.timeout})
	if opts.cache {
		rpc = service.NewCache(rpc)
	}

	return &client{
		rpc:       rpc,
		isMainnet: opts.isMainnet,
	}
}

func (c *client) argsFor(p positionalOptionalArgs, args ...interface{}) []interface{} {
	if reflect.ValueOf(p).IsNil() {
		return args
	}

	return append(args, p.Args()...)
}
