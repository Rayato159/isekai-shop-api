package config

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

var once sync.Once

type (
	DatabaseConfig struct {
		Host     string `mapstructure:"host" validate:"required"`
		Port     int    `mapstructure:"port" validate:"required"`
		User     string `mapstructure:"user" validate:"required"`
		Password string `mapstructure:"password" validate:"required"`
		DBName   string `mapstructure:"dbname" validate:"required"`
		SSLMode  string `mapstructure:"sslmode" validate:"required"`
		Schema   string `mapstructure:"schema" validate:"required"`
	}

	ServerConfig struct {
		Port         int           `mapstructure:"port" validate:"required"`
		AllowOrigins []string      `mapstructure:"allowOrigins" validate:"required"`
		Timeout      time.Duration `mapstructure:"timeout" validate:"required"`
		BodyLimit    string        `mapstructure:"bodyLimit" validate:"required"`
	}

	OAuth2Config struct {
		PlayerRedirectUrl string           `mapstructure:"playerRedirectUrl" validate:"required"`
		AdminRedirectUrl  string           `mapstructure:"adminRedirectUrl" validate:"required"`
		ClientId          string           `mapstructure:"clientId" validate:"required"`
		ClientSecret      string           `mapstructure:"clientSecret" validate:"required"`
		Endpoints         *oauth2Endpoints `mapstructure:"endpoints" validate:"required" `
		Scopes            []string         `mapstructure:"scopes" validate:"required"` // https://developers.google.com/identity/protocols/oauth2/scopes
		UserInfoUrl       string           `mapstructure:"userInfoUrl" validate:"required"`
		RevokeUrl         string           `mapstructure:"revokeUrl" validate:"required"`
	}

	oauth2Endpoints struct {
		AuthUrl       string `mapstructure:"authUrl" validate:"required"`
		TokenUrl      string `mapstructure:"tokenUrl" validate:"required"`
		DeviceAuthUrl string `mapstructure:"deviceAuthUrl" validate:"required"`
	}

	StateConfig struct {
		Secret    string        `mapstructure:"secret" validate:"required"`
		ExpiresAt time.Duration `mapstructure:"expiresAt" validate:"required"`
		Issuer    string        `mapstructure:"issuer" validate:"required"`
	}

	AppConfig struct {
		DatabaseConfig *DatabaseConfig `mapstructure:"database" validate:"required"`
		ServerConfig   *ServerConfig   `mapstructure:"server" validate:"required"`
		OAuth2Config   *OAuth2Config   `mapstructure:"oauth2" validate:"required"`
		StateConfig    *StateConfig    `mapstructure:"state" validate:"required"`
	}
)

var appConfigInstance *AppConfig

func GetAppConfig() *AppConfig {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config/")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := viper.ReadInConfig(); err != nil {
			panic(fmt.Errorf("read config file failed: %v", err))
		}

		if err := viper.Unmarshal(&appConfigInstance); err != nil {
			panic(fmt.Errorf("unmarshalkey config file failed: %v", err))
		}

		validate := validator.New()

		if err := validate.Struct(appConfigInstance); err != nil {
			panic(fmt.Errorf("validate config file failed: %v", err))
		}
	})

	return appConfigInstance
}
