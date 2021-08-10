package web

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/nch-bowstave/paymail"
	"github.com/nch-bowstave/paymail/config"
	"github.com/stretchr/testify/assert"
)

func TestCapability(t *testing.T) {
	e := echo.New()
	cfg := &config.Paymail{
		Domain:  "somedomain.com",
		Version: "1.0",
	}

	capability := &paymail.Capability{
		BSVAlias: "1.0",
		Capabilities: paymail.Capabilities{
			PKI:                "https://" + cfg.Domain + version + "/{alias}@{domain.tld}/id",
			PaymentDestination: "https://" + cfg.Domain + version + "/{alias}@{domain.tld}/payment-destination",
			PublicProfile:      "https://" + cfg.Domain + version + "/{alias}@{domain.tld}/public-profile",
			VerifyPublicKey:    "https://" + cfg.Domain + version + "/{alias}@{domain.tld}/verify-pubkey/{pubkey}",
		},
	}

	c := NewCapability(cfg)
	c.RegisterRoutes(e.Group(""))

	req := httptest.NewRequest(http.MethodGet, "/.well-known/bsvalias", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	t.Run("should return capability json", func(t *testing.T) {
		if assert.NoError(t, c.Capability(ctx)) {
			var body *paymail.Capability
			if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
				panic(err)
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, capability, body)
		}
	})

}
