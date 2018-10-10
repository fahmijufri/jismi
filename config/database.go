package config

type DatabaseConfig struct {
	databaseUsername string
	databasePassword string
	databaseAddress  string
	databaseName     string
	databaseLog      bool
	databaseMaxOpen  int
	databaseMaxIdle  int
}

func (dc DatabaseConfig) Username() string {
	return dc.databaseUsername
}

func (dc DatabaseConfig) Password() string {
	return dc.databasePassword
}

func (dc DatabaseConfig) Address() string {
	return dc.databaseAddress
}

func (dc DatabaseConfig) DatabaseName() string {
	return dc.databaseName
}

func (dc DatabaseConfig) LogEnabled() bool {
	return dc.databaseLog
}

func (dc DatabaseConfig) DatabaseMaxOpenConn() int {
	return dc.databaseMaxOpen
}

func (dc DatabaseConfig) DatabaseMaxIdleConn() int {
	return dc.databaseMaxIdle
}

func (dc DatabaseConfig) SetAddress(address string) {
	dc.databaseAddress = address
}
