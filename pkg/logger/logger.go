package logger

import (
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger interface {
	Debug(msg string)
	Debugf(format string, args ...any)
	Info(msg string)
	Infof(format string, args ...any)
	Error(err error, msg string)
	Errorf(err error, format string, args ...any)
	Errors(msg, key string, errs ...error)
	Errorsf(key string, errs []error, format string, args ...any)
	Fatal(err error, msg string)
	Fatalf(err error, format string, args ...any)
	WithStr(key, value string) Logger
}

func (l logger) Debug(msg string) {
	l.Logger.Debug().Msg(msg)
}

func (l logger) Debugf(format string, args ...any) {
	l.Logger.Debug().Msgf(format, args...)
}

func (l logger) Info(msg string) {
	l.Logger.Info().Msg(msg)
}

func (l logger) Infof(format string, args ...any) {
	l.Logger.Info().Msgf(format, args...)
}

func (l logger) Error(err error, msg string) {
	l.Logger.Error().Err(err).Msg(msg)
}

func (l logger) Errorf(err error, format string, args ...any) {
	l.Logger.Error().Err(err).Msgf(format, args...)
}

func (l logger) Errors(msg, key string, errs ...error) {
	l.Logger.Error().Errs(key, errs).Msg(msg)
}

func (l logger) Errorsf(key string, errs []error, format string, args ...any) {
	l.Logger.Error().Errs(key, errs).Msgf(format, args...)
}

func (l logger) Fatal(err error, msg string) {
	l.Logger.Fatal().Err(err).Msg(msg)
}

func (l logger) Fatalf(err error, format string, args ...any) {
	l.Logger.Fatal().Err(err).Msgf(format, args...)
}

func (l logger) WithStr(key, value string) Logger {
	log := l.Logger.With().Str(key, value).Logger()
	return logger{&log}
}

type logger struct {
	*zerolog.Logger
}

func New(debugMode bool, logDir string) (Logger, error) {
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		return nil, err
	}

	multiWriter := zerolog.MultiLevelWriter(
		createLevelConsoleWriter(),
		&lumberjack.Logger{
			Filename:   filepath.Join(logDir, "messenger_server.log"),
			MaxSize:    100,
			MaxBackups: 3,
			MaxAge:     30,
			Compress:   true,
		})

	zerolog.TimeFieldFormat = time.RFC3339

	level := zerolog.InfoLevel
	if debugMode {
		level = zerolog.DebugLevel
	}

	zerolog.SetGlobalLevel(level)

	log := zerolog.New(multiWriter).
		With().
		Str("logger", "messenger").
		Timestamp().
		Logger()

	if debugMode {
		dbgInfoLogger := log.With().Caller().Logger()
		dbgInfoLogger.Info().Msg("DEBUG MODE")
	}

	log = log.With().CallerWithSkipFrameCount(3).Logger()

	return logger{&log}, nil
}

type levelConsoleWriter struct {
	stdout zerolog.ConsoleWriter
	stderr zerolog.ConsoleWriter
}

func (w levelConsoleWriter) WriteLevel(level zerolog.Level, p []byte) (n int, err error) {
	if level <= zerolog.InfoLevel {
		return w.stdout.Write(p)
	}
	return w.stderr.Write(p)
}

func (w levelConsoleWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}

func createLevelConsoleWriter() levelConsoleWriter {
	writer := zerolog.ConsoleWriter{
		TimeFormat:    time.RFC3339,
		PartsOrder:    []string{"time", "logger", "caller", "level", "message"},
		FieldsExclude: []string{"logger"},
	}

	stdoutWrite := writer
	stdoutWrite.Out = os.Stdout

	stderrWriter := writer
	stderrWriter.Out = os.Stderr

	return levelConsoleWriter{
		stdout: stdoutWrite,
		stderr: stderrWriter,
	}
}
