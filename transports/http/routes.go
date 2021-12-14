package web

const (
	version                         = "/api/v2"
	routeCapability                 = "/.well-known/bsvalias"
	routePki                        = version + "/pki/:handle"
	routePaymentDestination         = version + "/:handle/payment-destination"          //	Returns an output to send money to a given paymail owner
	routePaymentDestinationResponse = version + "/:handle/payment-destination-response" //	Returns response to payment destination
	routePublicProfile              = version + "/:handle/public-profile"               //	Returns public key for a given paymail
	routeVerify                     = version + "/:handle/verify-pubkey/:pubkey"        // Returns bool
	routeRegister                   = version + "/register"
)
