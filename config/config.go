package config

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig
	Logger Logger
}

type ServerConfig struct {
	AppVersion   string
	Port         string
	Env          string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Logger struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

// LoadConfig Load the configuration file from given path
func LoadConfig() (*viper.Viper, error) {
	filename := getConfigPath()

	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		var notFoundError viper.ConfigFileNotFoundError
		if ok := errors.Is(err, notFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

// ParseConfig Parse the viper generated struct into the Config struct
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}

func getConfigPath() string {
	path := os.Getenv("config")
	if path == "docker" {
		return "./config/config-docker"
	}
	return "./config/config-local"
}
