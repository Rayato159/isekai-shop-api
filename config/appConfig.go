package config

import (
	"fmt"
	"sync"
	"time"

	"github.com/spf13/viper"
)

var once sync.Once

type (
	AppConfig struct {
		ServerConfig   *ServerConfig
		DatabaseConfig *DatabaseConfig
	}

	DatabaseConfig struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
		SSLMode  string
		Schema   string
	}

	ServerConfig struct {
		Port         int
		AllowOrigins []string
		Timeout      time.Duration
		BodyLimit    string
	}
)

var appConfig *AppConfig

func GetAppConfig() *AppConfig {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./")

		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("fatal error config file: %v", err))
		}

		appConfig = &AppConfig{
			ServerConfig:   getServerConfig(),
			DatabaseConfig: getDatabaseConfig(),
		}
	})

	return appConfig
}

func getServerConfig() *ServerConfig {
	return &ServerConfig{
		Port:         viper.GetInt("server.port"),
		AllowOrigins: viper.GetStringSlice("server.allowOrigins"),
		Timeout:      viper.GetDuration("server.timeout"),
		BodyLimit:    viper.GetString("server.bodyLimit"),
	}
}

func getDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetInt("database.port"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		DBName:   viper.GetString("database.dbname"),
		SSLMode:  viper.GetString("database.sslmode"),
		Schema:   viper.GetString("database.schema"),
	}
}
