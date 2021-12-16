package config

// Environment variable constants.
const (
	EnvPaymailVersion = "env.paymail.version"

	EnvPaymailPort = "paymail.port"
	EnvPaymailHost = "paymail.host"
	EnvPaydPort    = "payd.port"
	EnvPaydHost    = "payd.host"
	EnvP4Port      = "p4.port"
	EnvP4Host      = "p4.host"
	EnvNetwork     = "env.network"

	LogDebug = "debug"
	LogInfo  = "info"
	LogError = "error"
	LogWarn  = "warn"
)

type (
	// Config returns strongly typed config values.
	Config struct {
		Paymail *Paymail
		PayD    *PayD
		P4      *P4
	}

	// Paymail contains the version and url of the current paymail server.
	Paymail struct {
		Domain   string
		Port     string
		Hostname string
	}

	PayD struct {
		Port     string
		Hostname string
	}

	P4 struct {
		Port     string
		Hostname string
	}
)
