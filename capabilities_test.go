package paymail

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCapability(t *testing.T) {
	capabilityDocument, err := GenerateCapabilitiesDocument()
	assert.NoError(t, err, "must be no errors when generating dynamic capabilities file")
	t.Run("should return capability json", func(t *testing.T) {
		responseBodyBytes, err := json.Marshal(capabilityDocument)
		assert.NoError(t, err, "response marshalling failed")
		fmt.Println(string(responseBodyBytes))
	})
}
