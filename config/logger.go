package config

import "github.com/Be1chenok/messenger/pkg/defaultEnv"

type Logger struct {
	DebugMode bool
	LogDir    string
}

func newLogger() (Logger, error) {
	debugMode, err := defaultEnv.GetBool("DEBUG_MODE", false)
	if err != nil {
		return Logger{}, err
	}

	return Logger{
		DebugMode: debugMode,
		LogDir:    defaultEnv.GetString("LOG_DIR", "./logs"),
	}, nil
}
