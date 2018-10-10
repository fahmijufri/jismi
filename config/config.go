package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var config Config

type Config struct {
	serverAddress  string
	databaseConfig DatabaseConfig
	loggerFormat   string
	logLevel       string
}

func LoadConfig() {
	viper.SetDefault(Address, ":8080")
	viper.SetDefault(LoggerFormat, "text")

	viper.AutomaticEnv()

	viper.SetConfigName("application")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./../")
	viper.AddConfigPath("../../")
	viper.AddConfigPath("../../../")
	viper.AddConfigPath("../../../../")
	viper.SetConfigType("yaml")

	viper.ReadInConfig()

	config = Config{
		serverAddress: viper.GetString(Address),
		loggerFormat:  viper.GetString(LoggerFormat),
		logLevel:      viper.GetString(LogLevel),
		databaseConfig: DatabaseConfig{
			databaseUsername: getString(DatabaseUsername),
			databasePassword: getString(DatabasePassword),
			databaseAddress:  getString(DatabaseAddress),
			databaseName:     getString(DatabaseName),
			databaseLog:      getBool(DatabaseLog),
			databaseMaxOpen:  getInt(DatabaseMaxOpen),
			databaseMaxIdle:  getInt(DatabaseMaxIdle),
		},
	}
}

func Database() DatabaseConfig {
	return config.databaseConfig
}

func LogFormat() string {
	return config.loggerFormat
}

func ServerAddress() string {
	return config.serverAddress
}

func checkEnvKey(key string) {
	if !viper.IsSet(key) && os.Getenv(key) == "" {
		log.Fatalf("%v env key is not set", key)
	}
}

func getString(key string) string {
	checkEnvKey(key)
	return viper.GetString(key)
}

func getBool(key string) bool {
	checkEnvKey(key)
	return viper.GetBool(key)
}

func getInt(key string) int {
	checkEnvKey(key)
	return viper.GetInt(key)
}
