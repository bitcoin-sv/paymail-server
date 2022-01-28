package internal

import (
	"encoding/json"

	"github.com/nch-bowstave/paymail/config"
	"github.com/nch-bowstave/paymail/data"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// CapabilitiesDocument is the whole response body for the capability discovery mechanism of Paymail.
type CapabilitiesDocumentV1 struct {
	Version      string            `json:"bsvalias"`
	Capabilities map[string]string `json:"capabilities"`
}

// AddCapability is a function for dynamically adding capabilities from a yaml file.
func (caps *CapabilitiesDocumentV1) AddCapabilityV1(cfg *config.Paymail, d []byte) error {
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
		caps.Capabilities = map[string]string{}
	}
	caps.Capabilities[brfcID] = capability.Callback
	if brfcID == "0f9681ab58f0" {
		caps.Capabilities["pki"] = capability.Callback
	}
	return nil
}

func GenerateCapabilitiesDocumentV1(cfg *config.Paymail) error {
	var capabilities CapabilitiesDocumentV1
	capabilities.Version = "1.0"
	files, err := data.CapabilitiesData.LoadAll()
	if err != nil {
		return errors.Wrap(err, "failed to load capabilities")
	}
	for _, data := range files {
		if err = capabilities.AddCapabilityV1(cfg, data); err != nil {
			return errors.Wrapf(err, "failed to add capability %s", string(data))
		}
	}
	d, err := json.Marshal(capabilities)
	if err != nil {
		return errors.Wrap(err, "failed to marshal capabilities")
	}

	if err = data.OverwriteStaticCapabilitiesFileV1(d); err != nil {
		return errors.Wrap(err, "failed to overwrite existing capabilities file")
	}
	return nil
}
