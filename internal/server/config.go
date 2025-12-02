package server

import (
	"fmt"
	"time"
)

type Config struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	IdleTimeout time.Duration `mapstructure:"idle_timeout"`
	ReadTimeout time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
}

func (c *Config) Address() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}