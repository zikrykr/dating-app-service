package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/dating-app-service/config"
	"github.com/dating-app-service/constants"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GormDB struct {
	*gorm.DB
}

type dbConfig struct {
	GormDB       *GormDB
	ConnectionDB *sql.DB
}

func (db dbConfig) CloseConnection() error {
	return db.ConnectionDB.Close()
}

func Init() (dbConfig, error) {
	var (
		dbConfigVar dbConfig
		loggerGorm  logger.Interface
	)
	configData := config.GetConfig()

	loggerGorm = logger.Default.LogMode(logger.Silent)
	if configData.App.Env == constants.DEV {
		loggerGorm = logger.Default.LogMode(logger.Info)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", configData.DB.Host, configData.DB.User, configData.DB.Pass, configData.DB.Name),
	}), &gorm.Config{
		Logger:                 loggerGorm,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return dbConfigVar, err
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return dbConfigVar, err
	}

	sqlDB.SetConnMaxIdleTime(time.Second * time.Duration(configData.DB.MaxIdletimeConn))
	sqlDB.SetMaxIdleConns(configData.DB.MaxIdleConn)
	sqlDB.SetMaxOpenConns(configData.DB.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(configData.DB.MaxLifetimeConn))
	dbConfigVar.ConnectionDB = sqlDB

	dbConfigVar.GormDB = &GormDB{gormDB}
	logrus.Info("database connected")

	return dbConfigVar, nil
}
