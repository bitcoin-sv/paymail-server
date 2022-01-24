package main

import (
	"github.com/nch-bowstave/paymail/cmd/internal"
	"github.com/nch-bowstave/paymail/config"
)

func main() {
	// cfg := config.NewViperConfig("paymail").WithPaymail().Load() // not sure why this isn't working.
	cfg := &config.Config{
		Paymail: &config.Paymail{
			Root: "https://paymail.carefulbear.com",
		},
	}
	// generate a static capabilities document based on the files in data/capabilities.
	internal.GenerateCapabilitiesDocument(cfg)
	internal.GenerateCapabilitiesDocumentV1(cfg)
}
