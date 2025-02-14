package cmd

import (
	"errors"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

var Log *LoggerService

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Fatal(msg string)
}

type LoggerService struct {
	log zerolog.Logger
	env string
}

// TODO:
// 0. Add a logging middleware to log all requests
// 1. Handle log sampling
// 2. Handle log rotation
// 3. Create separate log files for everyday
// 4. Create a transport to a separate logging service
// 5. Colorize log-outputs using code from ./prettylogger.go
// 6. Improve prettified JSON-logging in console

func NewLoggerService(environment string, file *os.File) *LoggerService {
	var output io.Writer

	if environment == "development" {
		// Logging to both file and std.out during development
		consoleOut := zerolog.ConsoleWriter{Out: file, TimeFormat: time.RFC3339}
		fileOut := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
		output = zerolog.MultiLevelWriter(fileOut, consoleOut)
	} else if environment == "production" {
		// Logging only to std.out during production
		output = zerolog.ConsoleWriter{Out: file, TimeFormat: time.RFC3339}
	} else {
		panic(errors.New("Could not identify environment"))
	}

	// Default sampling setup
	logger := zerolog.New(output).With().Timestamp().Logger().Sample(zerolog.LevelSampler{
		InfoSampler: &zerolog.BasicSampler{N: 5},
	})
	return &LoggerService{
		log: logger,
		env: environment,
	}
}

func (l *LoggerService) Info(msg string) {
	l.log.WithLevel(zerolog.InfoLevel).Msgf("%s", msg)
}

func (l *LoggerService) Debug(msg string) {
	l.log.WithLevel(zerolog.DebugLevel).Msgf("%s", msg)
}

func (l *LoggerService) Warn(msg string) {
	l.log.WithLevel(zerolog.WarnLevel).Msgf("%s", msg)
}

func (l *LoggerService) Error(msg string) {
	l.log.WithLevel(zerolog.ErrorLevel).Msgf("%s", msg)
}

func (l *LoggerService) Fatal(msg string) {
	l.log.WithLevel(zerolog.FatalLevel).Msgf("%s", msg)
}
