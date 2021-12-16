package service

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
