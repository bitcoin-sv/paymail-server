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
	viper.SetDefault(EnvServerHost, "paymail")
	viper.SetDefault(EnvServerPort, ":8446")
	viper.SetDefault(EnvServerFQDN, "paymail:8446")
	viper.SetDefault(EnvServerSwaggerEnabled, true)
	viper.SetDefault(EnvServerSwaggerHost, "localhost:8446")

	// db
	viper.SetDefault(EnvDb, "sqlite")
	viper.SetDefault(EnvDbDsn, "file:data/wallet.db?_foreign_keys=true&pooled=true")
	viper.SetDefault(EnvDbSchema, "data/sqlite/migrations")
	viper.SetDefault(EnvDbMigrate, true)

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
