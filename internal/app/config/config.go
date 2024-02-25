package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

const (
	configType = "yaml"
	configName = "spec"

	// Env variable keys
	configPathEnvKey = "SPEC_FILE_PATH"
)

// Config represents the structured configuration settings.
type Config struct {
	App       App       `mapstructure:"app" validate:"required"`
	Database  Database  `mapstructure:"database" validate:"required"`
	Simulator Simulator `mapstructure:"simulator" validate:"required"`
}
type App struct {
	Name        string `mapstructure:"name" validate:"required"`
	Environment string `mapstructure:"environment" validate:"required"`
	Port        string `mapstructure:"port" validate:"required"`
}
type Simulator struct {
	UpdateCount  int `mapstructure:"updateCount" validate:"required"`
	SimulateTime int `mapstructure:"simulatedTime" validate:"required"`
}

type Database struct {
	Host     string `mapstructure:"host" validate:"required"`
	Port     string `mapstructure:"port" validate:"required"`
	Name     string `mapstructure:"name" validate:"required"`
	Username string `mapstructure:"username" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
}

// GetConfig retrieves the configuration settings.
func GetConfig() (*Config, error) {
	configPath := os.Getenv(configPathEnvKey)
	if configPath == "" {
		return nil, fmt.Errorf("config path Env variable is not set for key, %v", configPathEnvKey)
	}

	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	err := viper.ReadInConfig()

	if err != nil {
		return nil, fmt.Errorf("error while reading config file, %v", err)
	}

	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		return nil, fmt.Errorf("unable to decode into config struct, %v", err)
	}

	return config, nil
}
