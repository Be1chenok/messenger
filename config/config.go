package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

type Config struct {
	Logger     Logger
	HTTPServer HTTPServer
	Handler    Handler
}

func New(confPath string) (*Config, error) {
	if confPath != "" {
		if err := godotenv.Load(confPath); err != nil {
			return nil, fmt.Errorf("failed to load conf path: %w", err)
		}
	}

	logger, err := newLogger()
	if err != nil {
		return nil, err
	}

	httpServer, err := newHTTPServer()
	if err != nil {
		return nil, err
	}

	handler, err := newHandler()
	if err != nil {
		return nil, err
	}

	return &Config{
		Logger:     logger,
		HTTPServer: httpServer,
		Handler:    handler,
	}, nil
}
