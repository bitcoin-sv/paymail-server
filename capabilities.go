package paymail

import (
	"encoding/hex"
	"strings"

	"github.com/libsv/go-bt"
	"github.com/libsv/go-bt/crypto"
	"github.com/nch-bowstave/paymail/data"
	"gopkg.in/yaml.v2"
)

// CapabilitiesDocument is the whole response body for the capability discovery mechanism of Paymail.
type CapabilitiesDocument struct {
	Version      string                 `json:"version"`
	Capabilities map[string]*Capability `json:"capabilities"`
}

// Capability is a single capablity as defined by some brfcid specification.
type Capability struct {
	Title    string   `yaml:"title,omitempty" json:"omit"`
	Authors  []string `yaml:"authors,omitempty" json:"omit"`
	Version  string   `yaml:"version,omitempty" json:"omit"`
	Callback string   `yaml:"callback,omitempty" json:"callback,omitempty"`
}

// AddCapability is a function for dynamically adding capabilities from a yaml file.
func (caps *CapabilitiesDocument) AddCapability(d []byte) error {
	capability := &Capability{}
	err := yaml.Unmarshal(d, capability)
	if err != nil {
		return err
	}
	brfcID := GenerateBrfcID(capability)
	if caps.Capabilities == nil {
		caps.Capabilities = map[string]*Capability{}
	}
	caps.Capabilities[brfcID] = capability
	return nil
}

func GenerateBrfcID(c *Capability) string {
	cat := make([]byte, 0)
	if c.Title != "" {
		cat = append(cat, []byte(strings.TrimSpace(c.Title))...)
	}
	if c.Authors != nil && len(c.Authors) > 0 {
		for _, author := range c.Authors {
			cat = append(cat, []byte(strings.TrimSpace(author))...)
		}
	}
	if c.Version != "" {
		cat = append(cat, []byte(strings.TrimSpace(c.Version))...)
	}
	return hex.EncodeToString(bt.ReverseBytes(crypto.Sha256d(cat)[26:]))
}

func GenerateCapabilitiesDocument() (*CapabilitiesDocument, error) {
	var capabilities CapabilitiesDocument
	files, err := data.CapabilitiesData.LoadAll()
	if err != nil {
		return nil, err
	}
	for _, data := range files {
		err := capabilities.AddCapability(data)
		if err != nil {
			return nil, err
		}
	}
	return &capabilities, nil
}
