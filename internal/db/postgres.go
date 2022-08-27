package db

import (
	"media-service/internal/config"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func InitPostgresConn() error {
	DB, err = gorm.Open(postgres.Open(config.PostgresDSN()), &gorm.Config{})

	if err != nil {
		logrus.Error(err)
		return err
	}

	logrus.Info("Successfully connected to postgres database")

	return nil
}
