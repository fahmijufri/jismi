package main

import (
	"github.com/fahmijufri/jismi/config"
	"github.com/fahmijufri/jismi/infrastructure/provider/postgres"
	"github.com/fahmijufri/jismi/infrastructure/repository"
	"github.com/fahmijufri/jismi/infrastructure/web"
	"github.com/sirupsen/logrus"
	graceful "gopkg.in/tylerb/graceful.v1"
)

func init() {
	config.LoadConfig()
	config.SetupLogger()
}

func main() {
	db, err := postgres.ConnectSQL()
	if err != nil {
		logrus.WithError(err).Fatalln("Failed connecting to database")
	}
	defer postgres.CloseDB(db)

	server := &graceful.Server{
		Timeout: 0,
	}
	userRepository := repository.NewUserRepository(db)
	web.Run(server, userRepository)
}
