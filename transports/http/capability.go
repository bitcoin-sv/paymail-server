package web

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/nch-bowstave/paymail"
	"github.com/nch-bowstave/paymail/config"
)

type capability struct {
	C *paymail.Capability
}

// NewCapability returns an instance of capability.
func NewCapability(c *config.Paymail) *capability {
	return &capability{
		C: &paymail.Capability{
			BSVAlias: c.Version,
			Capabilities: paymail.Capabilities{
				PKI:                "https://" + c.Domain + version + "/{alias}@{domain.tld}/id",                  // returns public key for alias
				PaymentDestination: "https://" + c.Domain + version + "/{alias}@{domain.tld}/payment-destination", // returns output to send money to a given paymail owner
				PublicProfile:      "https://" + c.Domain + version + "/{alias}@{domain.tld}/public-profile",
				VerifyPublicKey:    "https://" + c.Domain + version + "/{alias}@{domain.tld}/verify-pubkey/{pubkey}", // checks if a given pubkey belongs to given paymail
			},
		},
	}
}

// RegisterRoutes is used to register routes with echo.
func (c *capability) RegisterRoutes(g *echo.Group) {
	g.GET(routeCapability, c.Capability)
}

// Capability returns json object containing capabilities.
func (c *capability) Capability(e echo.Context) error {
	return e.JSON(http.StatusOK, c.C)
}
