package service

import (
	"context"

	"github.com/libsv/p4-server/log"
	"github.com/nch-bowstave/paymail/data"
)

// CapabilitiesDocument is the whole response body for the capability discovery mechanism of Paymail.
type CapabilitiesDocument struct {
	Version      string                 `json:"version"`
	Capabilities map[string]*Capability `json:"capabilities"`
}

// CapabilitiesDocument is the whole response body for the capability discovery mechanism of Paymail.
type CapabilitiesDocumentV1 struct {
	Version      string            `json:"version"`
	Capabilities map[string]string `json:"capabilities"`
}

// Capability is a single capablity as defined by some brfcid specification.
type Capability struct {
	Title    string   `yaml:"title,omitempty" json:"-"`
	Authors  []string `yaml:"authors,omitempty" json:"-"`
	Version  string   `yaml:"version,omitempty" json:"-"`
	Callback string   `yaml:"callback,omitempty" json:"callback,omitempty"`
	Readme   string   `yaml:"readme,omitempty" json:"readme,omitempty"`
}

type paymail struct {
	l log.Logger
}

// NewPaymail will create and return a new paymail service.
func NewPaymail(l log.Logger) *paymail {
	return &paymail{
		l: l,
	}
}

// Paymail contains the handlers for paymail service endpoints.
type Paymail interface {
	Capabilities(ctx context.Context, url string) ([]byte, error)
}

// Capabilties will create the response from a static document and return it.
func (svc *paymail) Capabilities(ctx context.Context, path string) ([]byte, error) {
	if path == "/.well-known/bsvalias.json" {
		return data.CapabilitiesData.LoadStaticDocument()
	} else {
		return data.CapabilitiesData.LoadStaticDocumentV1()
	}
}
