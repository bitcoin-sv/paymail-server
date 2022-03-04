package config

import (
	"strings"

	"github.com/spf13/viper"
)

// ViperConfig contains viper based configuration data.
type ViperConfig struct {
	*Config
}

// NewViperConfig will setup and return a new viper based configuration handler.
func NewViperConfig(appname string) *ViperConfig {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	return &ViperConfig{
		Config: &Config{},
	}
}

// WithServer will setup the web server configuration if required.
func (v *ViperConfig) WithServer() ConfigurationLoader {
	v.Server = &Server{
		Port:           viper.GetString(EnvServerPort),
		Hostname:       viper.GetString(EnvServerHost),
		SwaggerEnabled: viper.GetBool(EnvServerSwaggerEnabled),
		SwaggerHost:    viper.GetString(EnvServerSwaggerHost),
		FQDN:           viper.GetString(EnvServerFQDN),
	}
	return v
}

// WithDeployment sets up the deployment configuration if required.
func (v *ViperConfig) WithDeployment(appName string) ConfigurationLoader {
	v.Deployment = &Deployment{
		Environment: viper.GetString(EnvEnvironment),
		Region:      viper.GetString(EnvRegion),
		Version:     viper.GetString(EnvVersion),
		Commit:      viper.GetString(EnvCommit),
		BuildDate:   viper.GetTime(EnvBuildDate),
		AppName:     appName,
	}
	return v
}

// WithLog sets up and returns log config.
func (v *ViperConfig) WithLog() ConfigurationLoader {
	v.Logging = &Logging{Level: viper.GetString(EnvLogLevel)}
	return v
}

// WithPayD sets up and returns PayD viper config.
func (v *ViperConfig) WithPayD() ConfigurationLoader {
	v.PayD = &PayD{
		Host:            viper.GetString(EnvPaydHost),
		Port:            viper.GetString(EnvPaydPort),
		Secure:          viper.GetBool(EnvPaydSecure),
		CertificatePath: viper.GetString(EnvPaydCertPath),
		Noop:            viper.GetBool(EnvPaydNoop),
	}
	return v
}

// WithSockets reads socket env vars.
func (v *ViperConfig) WithSockets() ConfigurationLoader {
	v.Sockets = &Socket{
		ChannelTimeout:  viper.GetDuration(EnvSocketChannelTimeoutSeconds),
		MaxMessageBytes: viper.GetInt(EnvSocketMaxMessageBytes),
	}
	return v
}

// WithTransports reads transport config.
func (v *ViperConfig) WithTransports() ConfigurationLoader {
	v.Transports = &Transports{
		Mode: viper.GetString(EnvTransportMode),
	}
	return v
}

// Load will return the underlying config setup.
func (v *ViperConfig) Load() *Config {
	return v.Config
}
