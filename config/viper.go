package config

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

// NewViperConfig will setup and return a new viper based configuration handler.
func NewViperConfig(appname string) *Config {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	return &Config{}
}

// WithServer will setup the web server configuration if required.
func (c *Config) WithServer() *Config {
	viper.SetDefault(EnvServerPort, ":4000")
	viper.SetDefault(EnvServerHost, "localhost:4000")
	c.Server = &Server{
		Port:     viper.GetString(EnvServerPort),
		Hostname: viper.GetString(EnvServerHost),
	}
	return c
}

// WithDeployment sets up the deployment configuration if required.
func (c *Config) WithDeployment(appName string) *Config {
	viper.SetDefault(EnvEnvironment, "dev")
	viper.SetDefault(EnvRegion, "test")
	viper.SetDefault(EnvCommit, "test")
	viper.SetDefault(EnvVersion, "test")
	viper.SetDefault(EnvBuildDate, time.Now().UTC())

	c.Deployment = &Deployment{
		Environment: viper.GetString(EnvEnvironment),
		Region:      viper.GetString(EnvRegion),
		Version:     viper.GetString(EnvVersion),
		Commit:      viper.GetString(EnvCommit),
		BuildDate:   viper.GetTime(EnvBuildDate),
		AppName:     appName,
	}
	return c
}

// WithLog sets up and returns log config.
func (c *Config) WithLog() *Config {
	viper.SetDefault(EnvLogLevel, "info")
	c.Logging = &Logging{Level: viper.GetString(EnvLogLevel)}
	return c
}

// WithDb sets up and returns database configuration.
func (c *Config) WithDb() *Config {
	viper.SetDefault(EnvDb, "postgres")
	viper.SetDefault(EnvDbDsn, "user=ps password=root dbname=paymail host=localhost sslmode=disable port=5432")
	viper.SetDefault(EnvDbSchema, "data/sql/migrations")
	viper.SetDefault(EnvDbMigrate, false)
	c.Db = &Db{
		Type:       DbType(viper.GetString(EnvDb)),
		Dsn:        viper.GetString(EnvDbDsn),
		SchemaPath: viper.GetString(EnvDbSchema),
		MigrateDb:  viper.GetBool(EnvDbMigrate),
	}
	return c
}

// WithCapability
func (c *Config) WithCapability() *Config {
	domain := viper.GetString(EnvServerDomain)
	viper.SetDefault(EnvBSVAliasVersion, "1.0")

	c.Paymail = &Paymail{
		Domain:  domain,
		Version: viper.GetString(EnvBSVAliasVersion),
	}

	return c
}
