package internal

import (
	"encoding/hex"
	"encoding/json"
	"strings"

	"github.com/bitcoin-sv/paymail/config"
	"github.com/bitcoin-sv/paymail/data"
	"github.com/libsv/go-bk/crypto"
	"github.com/libsv/go-bt/v2"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
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

// AddCapability is a function for dynamically adding capabilities from a yaml file.
func (caps *CapabilitiesDocument) AddCapability(cfg *config.Paymail, d []byte) error {
	capability := &Capability{}
	err := yaml.Unmarshal(d, capability)
	if err != nil {
		return err
	}
	if cfg != nil {
		capability.Callback = cfg.Root + capability.Callback
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
		for idx, author := range c.Authors {
			if idx > 0 {
				cat = append(cat, []byte(", ")...)
			}
			cat = append(cat, []byte(strings.TrimSpace(author))...)
		}
	}
	if c.Version != "" {
		cat = append(cat, []byte(strings.TrimSpace(c.Version))...)
	}
	return hex.EncodeToString(bt.ReverseBytes(crypto.Sha256d(cat)[26:]))
}

func GenerateCapabilitiesDocument(cfg *config.Paymail) error {
	var capabilities CapabilitiesDocument
	capabilities.Version = "2.0"
	files, err := data.CapabilitiesData.LoadAll()
	if err != nil {
		return errors.Wrap(err, "failed to load capabilities")
	}
	for _, data := range files {
		if err = capabilities.AddCapability(cfg, data); err != nil {
			return errors.Wrapf(err, "failed to add capability %s", string(data))
		}
	}
	d, err := json.Marshal(capabilities)
	if err != nil {
		return errors.Wrap(err, "failed to marshal capabilities")
	}
	if err = data.OverwriteStaticCapabilitiesFile(d); err != nil {
		return errors.Wrap(err, "failed to overwrite existing capabilities file")
	}
	return nil
}
