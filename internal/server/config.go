package server

import "time"

const (
	defaultHost = "localhost"
	defaultPort = ":8080"
	defaultIdleTimeout = time.Second * 120
	defaultReadTimeout = time.Second * 10
	defaultWriteTimeout = time.Second * 10
)

type Config struct {
	host string
	port string
	idleTimeout time.Duration
	readTimeout time.Duration
	writeTimeout time.Duration
}

func DefaultConfig() *Config {
	return &Config{
		host: defaultHost,
		port: defaultPort,
		idleTimeout: defaultIdleTimeout,
		readTimeout: defaultReadTimeout,
		writeTimeout: defaultWriteTimeout,
	}
}