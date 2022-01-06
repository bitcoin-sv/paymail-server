package internal

import (
	"encoding/json"
	"fmt"

	"github.com/nch-bowstave/paymail/config"
	"github.com/nch-bowstave/paymail/data"
	"gopkg.in/yaml.v2"
)

// CapabilitiesDocument is the whole response body for the capability discovery mechanism of Paymail.
type CapabilitiesDocumentV1 struct {
	Version      string            `json:"version"`
	Capabilities map[string]string `json:"capabilities"`
}

// AddCapability is a function for dynamically adding capabilities from a yaml file.
func (caps *CapabilitiesDocumentV1) AddCapabilityV1(cfg *config.Config, d []byte) error {
	capability := &Capability{}
	err := yaml.Unmarshal(d, capability)
	if err != nil {
		return err
	}
	if cfg != nil {
		capability.Callback = cfg.Paymail.Root + capability.Callback
	}
	brfcID := GenerateBrfcID(capability)
	if caps.Capabilities == nil {
		caps.Capabilities = map[string]string{}
	}
	caps.Capabilities[brfcID] = capability.Callback
	return nil
}

func GenerateCapabilitiesDocumentV1(cfg *config.Config) {
	var capabilities CapabilitiesDocumentV1
	capabilities.Version = "1.0"
	files, err := data.CapabilitiesData.LoadAll()
	if err != nil {
		panic(err)
	}
	for _, data := range files {
		err = capabilities.AddCapabilityV1(cfg, data)
		if err != nil {
			panic(err)
		}
	}
	d, err := json.Marshal(capabilities)
	if err != nil {
		panic(err)
	}
	err = data.OverwriteStaticCapabilitiesFileV1(d)
	if err != nil {
		fmt.Println(err)
	}
}
