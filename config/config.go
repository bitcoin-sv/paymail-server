package config

import (
	"fmt"
	"regexp"
	"time"

	validator "github.com/theflyingcodr/govalidator"
)

// Environment variable constants.
const (
	EnvBSVAliasVersion = "env.bsvalias.version"

	EnvServerPort   = "server.port"
	EnvServerHost   = "server.host"
	EnvServerDomain = "server.domain"
	EnvEnvironment  = "env.environment"
	EnvMainNet      = "env.mainnet"
	EnvRegion       = "env.region"
	EnvVersion      = "env.version"
	EnvCommit       = "env.commit"
	EnvBuildDate    = "env.builddate"
	EnvLogLevel     = "log.level"
	EnvDb           = "db.type"
	EnvDbSchema     = "db.schema.path"
	EnvDbDsn        = "db.dsn"
	EnvDbMigrate    = "db.migrate"

	LogDebug = "debug"
	LogInfo  = "info"
	LogError = "error"
	LogWarn  = "warn"
)

var reDbType = regexp.MustCompile(`mysql|postgres`)

// Supported database types.
const (
	DBMySQL    DbType = "mysql"
	DBPostgres DbType = "postgres"
)

type (
	// Config returns strongly typed config values.
	Config struct {
		Logging    *Logging
		Server     *Server
		Deployment *Deployment
		Db         *Db
		Paymail    *Paymail
	}

	// DbType is used to restrict the dbs we can support.
	DbType string

	// Db contains database information.
	Db struct {
		Type       DbType
		SchemaPath string
		Dsn        string
		MigrateDb  bool
	}

	// Deployment contains information relating to the current
	// deployed instance.
	Deployment struct {
		Environment string
		AppName     string
		Region      string
		Version     string
		Commit      string
		BuildDate   time.Time
	}

	// Paymail contains the version and url of the current paymail server.
	Paymail struct {
		Domain  string
		Version string
	}

	// Logging contains log configuration.
	Logging struct {
		Level string
	}

	// Server contains all settings required to run a web server.
	Server struct {
		Port     string
		Hostname string
	}
)

// Validate will check config values are valid and return a list of failures
// if any have been found.
func (c *Config) Validate() error {
	vl := validator.New()
	if c.Db != nil {
		vl = vl.Validate("db.type", validator.MatchString(string(c.Db.Type), reDbType))
	}
	return vl.Err()
}

// IsDev determines if this app is running on a dev environment.
func (d *Deployment) IsDev() bool {
	return d.Environment == "dev"
}

func (d *Deployment) String() string {
	return fmt.Sprintf("Environment: %s \n AppName: %s\n Region: %s\n Version: %s\n Commit:%s\n BuildDate: %s\n",
		d.Environment, d.AppName, d.Region, d.Version, d.Commit, d.BuildDate)
}
