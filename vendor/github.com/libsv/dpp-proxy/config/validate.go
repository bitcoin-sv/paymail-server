package config

import validator "github.com/theflyingcodr/govalidator"

// Validate the configuration.
func (c *Config) Validate() error {
	v := validator.New()
	if c.Transports != nil {
		v = v.Validate("transport.mode", validator.AnyString(c.Transports.Mode, TransportModeHTTP, TransportModeHybrid, TransportModeSocket))
	}

	return v.Err()
}
