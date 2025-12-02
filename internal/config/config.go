package config

import (
	"fmt"

	"github.com/YotoHana/hitalent-test-case/internal/database"
	"github.com/YotoHana/hitalent-test-case/internal/server"
	"github.com/spf13/viper"
)

const (
	defaultConfigPath = "./"
)

type Config struct {
	Server   server.Config   `mapstructure:"server"`
	Database database.Config `mapstructure:"database"`
}

func Load(configPath string) (*Config, error) {
	v := viper.New()

	if configPath != "" {
		v.AddConfigPath(configPath)
	} else {
		v.AddConfigPath(defaultConfigPath)
	}

	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config: %v", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %v", err)
	}

	return &cfg, nil
}
