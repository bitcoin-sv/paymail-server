package config

import (
	"fmt"
	"time"
)

// Environment variable constants.
const (
	EnvServerPort                  = "server.port"
	EnvServerHost                  = "server.host"
	EnvServerFQDN                  = "server.fqdn"
	EnvServerSwaggerEnabled        = "server.swagger.enabled"
	EnvServerSwaggerHost           = "server.swagger.host"
	EnvEnvironment                 = "env.environment"
	EnvRegion                      = "env.region"
	EnvVersion                     = "env.version"
	EnvCommit                      = "env.commit"
	EnvBuildDate                   = "env.builddate"
	EnvLogLevel                    = "log.level"
	EnvPaydHost                    = "payd.host"
	EnvPaydPort                    = "payd.port"
	EnvPaydSecure                  = "payd.secure"
	EnvPaydCertPath                = "payd.cert.path"
	EnvPaydNoop                    = "payd.noop"
	EnvDPPTimeout                  = "dpp.timeout"
	EnvDPPHost                     = "dpp.host"
	EnvSocketChannelTimeoutSeconds = "socket.channel.timeoutseconds"
	EnvSocketMaxMessageBytes       = "socket.maxmessage.bytes"
	EnvTransportMode               = "transport.mode"
	EnvPaymailRoot                 = "paymail.root"
	EnvDb                          = "db.type"
	EnvDbSchema                    = "db.schema.path"
	EnvDbDsn                       = "db.dsn"
	EnvDbMigrate                   = "db.migrate"

	LogDebug = "debug"
	LogInfo  = "info"
	LogError = "error"
	LogWarn  = "warn"

	TransportModeHybrid = "hybrid"
	TransportModeHTTP   = "http"
	TransportModeSocket = "socket"
)

// Config returns strongly typed config values.
type Config struct {
	Logging    *Logging
	Db         *Db
	Server     *Server
	Deployment *Deployment
	PayD       *PayD
	DPP        *DPP
	Paymail    *Paymail
	Sockets    *Socket
	Transports *Transports
}

// Deployment contains information relating to the current
// deployed instance.
type Deployment struct {
	Environment string
	AppName     string
	Region      string
	Version     string
	Commit      string
	BuildDate   time.Time
}

// Db contains database information.
type Db struct {
	Type       DbType
	SchemaPath string
	Dsn        string
	MigrateDb  bool
}

// DbType is used to restrict the dbs we can support.
type DbType string

// Supported database types.
const (
	DBSqlite   DbType = "sqlite"
	DBMySQL    DbType = "mysql"
	DBPostgres DbType = "postgres"
)

// IsDev determines if this app is running on a dev environment.
func (d *Deployment) IsDev() bool {
	return d.Environment == "dev"
}

func (d *Deployment) String() string {
	return fmt.Sprintf("Environment: %s \n AppName: %s\n Region: %s\n Version: %s\n Commit:%s\n BuildDate: %s\n",
		d.Environment, d.AppName, d.Region, d.Version, d.Commit, d.BuildDate)
}

// Logging contains log configuration.
type Logging struct {
	Level string
}

// Server contains all settings required to run a web server.
type Server struct {
	Port     string
	Hostname string
	// FQDN - fully qualified domain name, used to form the paymentRequest
	// payment URL as this may be different from the hostname + port.
	FQDN string
	// SwaggerEnabled if true we will include an endpoint to serve swagger documents.
	SwaggerEnabled bool
	SwaggerHost    string
}

// PayD is used to setup connection to a payd instance.
// In this case, we connect to only one merchant wallet
// implementors may need to connect to more.
type PayD struct {
	Host            string
	Port            string
	Secure          bool
	CertificatePath string
	Noop            bool
}

// DPP contains information relating to a dpp interactions.
type DPP struct {
	Timeout    int
	ServerHost string
}

type Paymail struct {
	Root string
}

// Socket contains config items for a socket server.
type Socket struct {
	MaxMessageBytes int
	ChannelTimeout  time.Duration
}

// Transports enables or disables dpp transports.
type Transports struct {
	Mode string
}

// ConfigurationLoader will load configuration items
// into a struct that contains a configuration.
type ConfigurationLoader interface {
	WithServer() ConfigurationLoader
	WithDb() ConfigurationLoader
	WithDeployment(app string) ConfigurationLoader
	WithLog() ConfigurationLoader
	WithPayD() ConfigurationLoader
	WithDPP() ConfigurationLoader
	WithSockets() ConfigurationLoader
	WithTransports() ConfigurationLoader
	WithPaymail() ConfigurationLoader
	Load() *Config
}
