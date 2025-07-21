package config

import (
	"fmt"
	"time"

	"github.com/Be1chenok/messenger/pkg/defaultEnv"
)

type Handler struct {
	RequestTimeout time.Duration
}

func newHandler() (Handler, error) {
	requestTimeout, err := defaultEnv.GetDuration("HANDLER_REQUEST_TIMEOUT", 5*time.Second)
	if err != nil {
		return Handler{}, err
	}
	if requestTimeout < 0 {
		return Handler{}, fmt.Errorf("handler request timeout: %w", ErrNegativeTimeout)
	}

	return Handler{
		RequestTimeout: requestTimeout,
	}, nil
}
