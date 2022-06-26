package configs

import (
	"fmt"
	"time"

	"github.com/creasty/defaults"
	"github.com/spf13/viper"
)

type Environment string

const (
	DEVELOPMENT Environment = "development"
	PRODUCTION  Environment = "production"
)

type AppConfig struct {
	STAGE Environment `default:"development"`
	// Define server settings.
	SERVER struct {
		HOST         string        `default:"0.0.0.0" mapstructure:"SERVER_HOST"`
		PORT         int           `default:"3000" mapstructure:"SERVER_PORT"`
		READ_TIMEOUT time.Duration `default:"10s" mapstructure:"SERVER_READ_TIMEOUT"`
		PREFORK      bool          `default:"false" mapstructure:"SERVER_PREFORK"`
	} `mapstructure:",squash"`
	// Define cache settings.
	JWT struct {
		SECRET_KEY         string        `mapstructure:"JWT_SECRET_KEY"`
		SECRET_KEY_EXPIRE  time.Duration `default:"10m" mapstructure:"JWT_SECRET_KEY_EXPIRE"`
		REFRESH_KEY        string        `mapstructure:"JWT_REFRESH_KEY"`
		REFRESH_KEY_EXPIRE time.Duration `default:"1M" mapstructure:"JWT_REFRESH_KEY_EXPIRE"`
	} `mapstructure:",squash"`
	// Define database settings.
	DB struct {
		HOST          string `default:"localhost" mapstructure:"DB_HOST"`
		PORT          string `default:"5432" mapstructure:"DB_PORT"`
		NAME          string `mapstructure:"DB_NAME"`
		USERNAME      string `mapstructure:"DB_USERNAME"`
		PASSWORD      string `mapstructure:"DB_PASSWORD"`
		SSL           string `default:"false" mapstructure:"DB_SSL"`
		MAX_CONN      int    `default:"100" mapstructure:"DB_MAX_CONN"`
		MAX_IDLE_CONN int    `default:"10" mapstructure:"DB_MAX_IDLE_CONN"`
	} `mapstructure:",squash"`
	// Define cache settings.
	CACHE struct {
		HOST      string `default:"localhost" mapstructure:"CACHE_HOST"`
		PORT      string `default:"5432" mapstructure:"CACHE_PORT"`
		USERNAME  string `mapstructure:"CACHE_USERNAME"`
		PASSWORD  string `mapstructure:"CACHE_PASSWORD"`
		SSL       bool   `default:"false" mapstructure:"CACHE_SSL"`
		DB_NUMBER int    `default:"1" mapstructure:"CACHE_DB_NUMBER"`
	} `mapstructure:",squash"`
}

func (c *AppConfig) SetDefaults() {
}

var cfg *AppConfig

func load() {
	newCfg := &AppConfig{}
	if err := defaults.Set(newCfg); err != nil {
		panic("Config Default Fail:" + err.Error())
	}
	if err := viper.Unmarshal(&newCfg); err != nil {
		panic("Config Initialize Fail:" + err.Error())
	}
	cfg = newCfg
}
func setup(path, configName string) {
	viper.SetConfigType("env")
	viper.SetConfigName(configName)
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			println("Not found config file", err.Error())
			// Config file not found; ignore error if desired
		} else {
			// Config file was found but another error was produced
			panic("Config Load Fail:" + err.Error())
		}
	}
	fmt.Println("File path:", viper.ConfigFileUsed())
}

func Get() AppConfig {
	if cfg == nil {
		setup(".", "app")
		load()
	}
	return *cfg
}
