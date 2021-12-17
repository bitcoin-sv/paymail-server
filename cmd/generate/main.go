package main

import (
	"github.com/nch-bowstave/paymail/cmd/internal"
)

func main() {
	// generate a static capabilities document based on the files in data/capabilities.
	internal.GenerateCapabilitiesDocument()
}
