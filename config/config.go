package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

type Config struct {
	HTTPServer HTTPServer
}

func New(confPath string) (*Config, error) {
	if confPath != "" {
		if err := godotenv.Load(confPath); err != nil {
			return nil, fmt.Errorf("failed to load conf path: %w", err)
		}
	}

	httpServer, err := newHTTPServer()
	if err != nil {
		return nil, err
	}

	return &Config{
		HTTPServer: httpServer,
	}, nil
}
