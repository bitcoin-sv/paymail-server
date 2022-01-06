package main

import (
	"github.com/nch-bowstave/paymail/cmd/internal"
	"github.com/nch-bowstave/paymail/config"
)

func main() {
	// manually load prefix to endpoints if required
	cfg := &config.Config{
		Paymail: &config.Paymail{
			Root: "http://localhost:8446",
		},
	}
	// generate a static capabilities document based on the files in data/capabilities.
	internal.GenerateCapabilitiesDocument(cfg)
	internal.GenerateCapabilitiesDocumentV1(cfg)
}
