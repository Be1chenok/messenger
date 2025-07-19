package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

const logFileName = "./logs/platform.log"

func New() (*zerolog.Logger, error) {
	if err := os.MkdirAll("./logs", os.ModePerm); err != nil {
		return nil, err
	}

	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}

	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	multiWriter := zerolog.MultiLevelWriter(consoleWriter, file)

	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	log := zerolog.New(multiWriter).
		With().
		Timestamp().
		Logger()

	return &log, nil
}
