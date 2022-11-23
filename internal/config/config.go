package config

import (
	"fmt"
	"os"
)

type Config struct {
	dbUser string
	dbPswd string
	dbHost string
	dbPort string
	dbName string
}

func Get() *Config {

	return &Config{
		dbUser: os.Getenv("DB_USER"),
		dbPswd: os.Getenv("DB_PASSWORD"),
		dbHost: os.Getenv("DB_HOST"),
		dbPort: os.Getenv("DB_PORT"),
		dbName: os.Getenv("DB_NAME"),
	}

}

func (c *Config) GetUrlConnector() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.dbUser,
		c.dbPswd,
		c.dbHost,
		c.dbPort,
		c.dbName,
	)
}
