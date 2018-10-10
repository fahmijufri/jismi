package config

const (
	Version = "0.0.1"
	AppName = "Jismi"

	// Environment variable name
	Address      = "ADDRESS"
	LoggerFormat = "LOGGER_FORMAT"
	LogLevel     = "LOG_LEVEL"

	// Database environment variable
	DatabaseUsername = "DATABASE_USERNAME"
	DatabasePassword = "DATABASE_PASSWORD"
	DatabaseAddress  = "DATABASE_HOST"
	DatabaseName     = "DATABASE_NAME"
	DatabaseLog      = "DATABASE_LOG"
	DatabaseMaxOpen  = "DATABASE_MAX_OPEN_CONN"
	DatabaseMaxIdle  = "DATABASE_MAX_IDLE_CONN"

	// Healthcheck possible path
	PathLiveness  = "/liveness"
	PathReadiness = "/readiness"
)
