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
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"dbname"`
		SSLMode  string `mapstructure:"sslmode"`
		Schema   string `mapstructure:"schema"`
	}

	ServerConfig struct {
		Port         int           `mapstructure:"port"`
		AllowOrigins []string      `mapstructure:"allowOrigins"`
		Timeout      time.Duration `mapstructure:"timeout"`
		BodyLimit    string        `mapstructure:"bodyLimit"`
	}

	OAuth2Config struct {
		ClientId     string           `mapstructure:"clientId"`
		ClientSecret string           `mapstructure:"clientSecret"`
		RedirectUrl  string           `mapstructure:"redirectUrl"`
		Endpoints    *oauth2Endpoints `mapstructure:"endpoints"`
		Scopes       []string         `mapstructure:"scopes"` // https://developers.google.com/identity/protocols/oauth2/scopes
		UserInfoUrl  string           `mapstructure:"userInfoUrl"`
		RevokeUrl    string           `mapstructure:"revokeUrl"`
	}

	oauth2Endpoints struct {
		AuthUrl       string `mapstructure:"authUrl"`
		TokenUrl      string `mapstructure:"tokenUrl"`
		DeviceAuthUrl string `mapstructure:"deviceAuthUrl"`
	}

	StateConfig struct {
		Secret    string        `mapstructure:"secret"`
		ExpiresAt time.Duration `mapstructure:"expiresAt"`
		Issuer    string        `mapstructure:"issuer"`
	}

	AppConfig struct {
		DatabaseConfig *DatabaseConfig `mapstructure:"database"`
		ServerConfig   *ServerConfig   `mapstructure:"server"`
		OAuth2Config   *OAuth2Config   `mapstructure:"oauth2"`
		StateConfig    *StateConfig    `mapstructure:"state"`
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
			panic(fmt.Errorf("read config file failed: %v", err))
		}

		if err := viper.UnmarshalKey("app", &appConfigInstance); err != nil {
			panic(fmt.Errorf("unmarshalkey config file failed: %v", err))
		}
	})

	return appConfigInstance
}
