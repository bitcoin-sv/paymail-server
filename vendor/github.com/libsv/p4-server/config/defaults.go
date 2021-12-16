package config

import (
	"time"

	"github.com/spf13/viper"
)

// SetupDefaults will set environment variables to default values.
//
// These can be overwritten when running the service.
func SetupDefaults() {
	// Web server defaults
	viper.SetDefault(EnvServerHost, "p4")
	viper.SetDefault(EnvServerPort, ":8445")
	viper.SetDefault(EnvServerFQDN, "p4:8445")
	viper.SetDefault(EnvServerSwaggerEnabled, true)
	viper.SetDefault(EnvServerSwaggerHost, "localhost:8445")

	// Environment Defaults
	viper.SetDefault(EnvEnvironment, "dev")
	viper.SetDefault(EnvRegion, "local")
	viper.SetDefault(EnvCommit, "test")
	viper.SetDefault(EnvVersion, "v0.0.0")
	viper.SetDefault(EnvBuildDate, time.Now().UTC())

	// Log level defaults
	viper.SetDefault(EnvLogLevel, "info")

	// PayD wallet Defaults
	viper.SetDefault(EnvPaydHost, "payd")
	viper.SetDefault(EnvPaydPort, ":8443")
	viper.SetDefault(EnvPaydSecure, false)
	viper.SetDefault(EnvPaydNoop, false)

	// Socket settings
	viper.SetDefault(EnvSocketChannelTimeoutSeconds, 7200*time.Second) // 2 hrs in seconds
	viper.SetDefault(EnvSocketMaxMessageBytes, 10000)

	// Transport settings
	viper.SetDefault(EnvTransportMode, TransportModeHTTP)
}
