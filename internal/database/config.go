package database

const (
	defaultHost = "localhost"
	defaultPort = ":5432"
	defaultUser = "dev"
	defaultPassword = "dev"
	defaultDBName = "dev"
	defaultSSLMode = "false"
)

type Config struct {
	host string
	port string
	user string
	password string
	dbName string
	sslMode string
}

func DefaultConfig() *Config {
	return &Config{
		host: defaultHost,
		port: defaultPort,
		user: defaultUser,
		password: defaultPassword,
		dbName: defaultDBName,
		sslMode: defaultSSLMode,
	}
}