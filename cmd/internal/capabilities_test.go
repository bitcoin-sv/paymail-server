package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCapability(t *testing.T) {
	t.Run("create a dynamically generated capabilities file", func(t *testing.T) {
		err := GenerateCapabilitiesDocument()
		assert.NoError(t, err, "must be no errors when generating dynamic capabilities file")
	})
}
