package service

import (
	"context"
	"encoding/json"

	"github.com/libsv/p4-server/log"
	"github.com/nch-bowstave/paymail/data"
)

// CapabilitiesDocument is the whole response body for the capability discovery mechanism of Paymail.
type CapabilitiesDocument struct {
	Version      string                 `json:"version"`
	Capabilities map[string]*Capability `json:"capabilities"`
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
	Capabilities(ctx context.Context) (*CapabilitiesDocument, error)
}

// Capabilties will create the response from a static document and return it.
func (svc *paymail) Capabilities(ctx context.Context) (*CapabilitiesDocument, error) {
	data, err := data.CapabilitiesData.LoadStaticDocument()
	if err != nil {
		svc.l.Error(err, "capabilities.json document not found")
		return nil, err
	}
	caps := &CapabilitiesDocument{}
	err = json.Unmarshal(data, caps)
	if err != nil {
		return nil, err
	}
	return caps, err
}
