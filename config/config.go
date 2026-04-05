
package config

import (
	"os"
)

type Config struct {
	ServerPort string
	LogLevel   string
}

func LoadConfig() (*Config, error) {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}

	return &Config{
		ServerPort: port,
		LogLevel:   logLevel,
	}, nil
}
