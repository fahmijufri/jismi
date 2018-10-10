package postgres

import (
	"errors"
	"fmt"

	"github.com/fahmijufri/jismi/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/matryer/resync"
	"github.com/sirupsen/logrus"
)

var (
	db      *gorm.DB
	err     error
	runOnce resync.Once
)

func ConnectSQL() (*gorm.DB, error) {
	databaseURL := fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable", config.Database().Username(),
		config.Database().Password(), config.Database().Address(), config.Database().DatabaseName())

	runOnce.Do(func() {
		db, err = gorm.Open("postgres", databaseURL)

		if err != nil {
			logrus.WithField("database_url", databaseURL).WithError(err).Errorln("Failed connect to database")

			return
		}

		db.DB().SetMaxIdleConns(config.Database().DatabaseMaxIdleConn())
		db.DB().SetMaxOpenConns(config.Database().DatabaseMaxOpenConn())

		err = db.DB().Ping()
		if err != nil {
			logrus.WithField("database_url", databaseURL).WithError(err).Errorln("Failed ping to database")

			return
		}

		db.LogMode(config.Database().LogEnabled())
	})

	return db, err
}

func SQLPing(db *gorm.DB) error {
	if db == nil {
		return errors.New("database object is nil")
	}
	return db.DB().Ping()
}

func CloseDB(db *gorm.DB) error {
	if db == nil {
		return errors.New("database object is nil")
	}
	return db.Close()
}

func ResetDBSingleton() {
	runOnce.Reset()
}
