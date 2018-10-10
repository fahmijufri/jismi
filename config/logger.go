package config

import (
	"os"
	"strings"

	"github.com/TV4/logrus-stackdriver-formatter"
	"github.com/sirupsen/logrus"
)

func SetupLogger() {
	logrus.SetOutput(os.Stderr)

	switch strings.ToLower(LogFormat()) {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	case "stackdriver":
		logrus.SetFormatter(stackdriver.NewFormatter(
			stackdriver.WithService(AppName),
			stackdriver.WithVersion(Version),
		))
	default:
		logrus.SetFormatter(&logrus.TextFormatter{})
	}
}
