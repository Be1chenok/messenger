package config

import (
	"fmt"
	"time"

	"github.com/Be1chenok/messenger/pkg/defaultEnv"
)

const (
	defaultMaxHeaderBytes int = 5 << 20
)

type HTTPServer struct {
	Port                  int
	ReadTimeout           time.Duration
	WriteTimeout          time.Duration
	MaxHeaderBytes        int
	ServerShutdownTimeout time.Duration
}

func newHTTPServer() (HTTPServer, error) {
	port, err := defaultEnv.GetInt("HTTP_SERVER_PORT", 8080)
	if err != nil {
		return HTTPServer{}, err
	}

	readTimeout, err := defaultEnv.GetDuration("HTTP_SERVER_READ_TIMEOUT", 20*time.Second)
	if err != nil {
		return HTTPServer{}, err
	}
	if readTimeout < 0 {
		return HTTPServer{}, fmt.Errorf("http server read timeout: %w", ErrNegativeTimeout)
	}

	writeTimeout, err := defaultEnv.GetDuration("HTTP_SERVER_WRITE_TIMEOUT", 20*time.Second)
	if err != nil {
		return HTTPServer{}, err
	}
	if writeTimeout < 0 {
		return HTTPServer{}, fmt.Errorf("http server write timeout: %w", ErrNegativeTimeout)
	}

	maxHeaderBytes, err := defaultEnv.GetInt("HTTP_SERVER_MAX_HEADER_BYTES", defaultMaxHeaderBytes)
	if err != nil {
		return HTTPServer{}, err
	}

	serverShutdownTimeout, err := defaultEnv.GetDuration("HTTP_SERVER_SHUTDOWN_TIMEOUT", 5*time.Minute)
	if err != nil {
		return HTTPServer{}, err
	}
	if serverShutdownTimeout < 0 {
		return HTTPServer{}, fmt.Errorf("http server shutdown timeout: %w", ErrNegativeTimeout)
	}

	return HTTPServer{
		Port:                  port,
		ReadTimeout:           readTimeout,
		WriteTimeout:          writeTimeout,
		MaxHeaderBytes:        maxHeaderBytes,
		ServerShutdownTimeout: serverShutdownTimeout,
	}, nil
}
