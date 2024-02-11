package config

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/spf13/viper"
)

var once sync.Once

type (
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

	OAuth2Config struct {
		ClientId     string
		ClientSecret string
		RedirectUrl  string
		Scopes       []string // https://developers.google.com/identity/protocols/oauth2/scopes
		UserInfoUrl  string
	}

	StateConfig struct {
		Secret    []byte
		ExpiresAt time.Duration
		Issuer    string
	}

	AppConfig struct {
		DatabaseConfig *DatabaseConfig
		ServerConfig   *ServerConfig
		OAuth2Config   *OAuth2Config
		StateConfig    *StateConfig
	}
)

var appConfigInstance *AppConfig

func GetAppConfig() *AppConfig {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("fatal error config file: %v", err))
		}

		appConfigInstance = &AppConfig{
			DatabaseConfig: getDatabaseConfig(),
			ServerConfig:   getServerConfig(),
			OAuth2Config:   getOAuth2Config(),
			StateConfig:    getStateConfig(),
		}
	})

	return appConfigInstance
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

func getServerConfig() *ServerConfig {
	return &ServerConfig{
		Port:         viper.GetInt("server.port"),
		AllowOrigins: viper.GetStringSlice("server.allowOrigins"),
		Timeout:      viper.GetDuration("server.timeout"),
		BodyLimit:    viper.GetString("server.bodyLimit"),
	}
}

func getOAuth2Config() *OAuth2Config {
	return &OAuth2Config{
		ClientId:     viper.GetString("oauth2.google.clientId"),
		ClientSecret: viper.GetString("oauth2.google.clientSecret"),
		RedirectUrl:  viper.GetString("oauth2.google.redirectUrl"),
		Scopes:       viper.GetStringSlice("oauth2.google.scopes"),
		UserInfoUrl:  viper.GetString("oauth2.google.userInfoUrl"),
	}
}

func getStateConfig() *StateConfig {
	return &StateConfig{
		Secret:    []byte(viper.GetString("state.jwt.secret")),
		ExpiresAt: viper.GetDuration("state.jwt.expiresAt"),
		Issuer:    viper.GetString("state.jwt.issuer"),
	}
}
