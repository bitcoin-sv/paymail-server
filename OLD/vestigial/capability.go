package paymail

import "context"

// ref: https://docs.moneybutton.com/docs/paymail/paymail-02-02-capability-discovery.html

type (
	// CapabilitService contains the handler for the capability endpoint.
	CapabilitService interface {
		Capability(ctx context.Context) (*Capability, error)
	}
)
