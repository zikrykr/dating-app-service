package config

import (
	"fmt"
	"log"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type (
	DB struct {
		Name            string
		User            string
		Pass            string
		Host            string
		MaxOpenConn     int
		MaxIdleConn     int
		MaxLifetimeConn int
		MaxIdletimeConn int
	}

	app struct {
		Env       string
		Name      string
		JWTSecret string
	}

	http struct {
		Port int
	}

	Config struct {
		DB   DB
		App  app
		Http http
	}
)

const (
	appName = "dating-app-service"
)

var (
	configData *Config
)

func InitConfig() {
	viper.SetConfigType("env")
	viper.SetConfigName(".env") // name of Config file (without extension)
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logrus.WithError(err).Warn("failed to load config file")
	}

	configData = &Config{
		DB: DB{
			Name:            getRequiredString("DB_NAME"),
			User:            getRequiredString("DB_USER"),
			Pass:            getRequiredString("DB_PASS"),
			Host:            getRequiredString("DB_HOST"),
			MaxOpenConn:     getRequiredInt("DB_MAX_OPEN_CONN"),
			MaxIdleConn:     getRequiredInt("DB_MAX_IDLE_CONN"),
			MaxLifetimeConn: getRequiredInt("DB_MAX_LIFETIME_CONN"),
			MaxIdletimeConn: getRequiredInt("DB_MAX_IDLETIME_CONN"),
		},
		App: app{
			Env:       getRequiredString("APP_ENV"),
			Name:      appName,
			JWTSecret: getRequiredString("APP_JWT_SECRET"),
		},
		Http: http{
			Port: getRequiredInt("APP_PORT"),
		},
	}
}

func getRequiredString(key string) string {
	if viper.IsSet(key) {
		return viper.GetString(key)
	}

	log.Fatalln(fmt.Errorf("KEY %s IS MISSING", key))
	return ""
}

func getRequiredInt(key string) int {
	if viper.IsSet(key) {
		return viper.GetInt(key)
	}

	panic(fmt.Errorf("KEY %s IS MISSING", key))
}

func getRequiredBool(key string) bool {
	if viper.IsSet(key) {
		return viper.GetBool(key)
	}

	panic(fmt.Errorf("KEY %s IS MISSING", key))
}

func getRequiredDuration(key string) time.Duration {
	if viper.IsSet(key) {
		return viper.GetDuration(key)
	}

	panic(fmt.Errorf("KEY %s IS MISSING", key))
}

func GetConfig() Config {
	return *configData
}
