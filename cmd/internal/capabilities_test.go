package internal

import (
	"testing"
)

func TestCapability(t *testing.T) {
	t.Run("create a dynamically generated capabilities file", func(t *testing.T) {
		GenerateCapabilitiesDocument()
	})
}
