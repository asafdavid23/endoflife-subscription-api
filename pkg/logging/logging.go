package logging

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func SetupLogger(level zerolog.Level) zerolog.Logger {
	// Set the global log level
	zerolog.SetGlobalLevel(level)

	// Set the time format for the logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// Create a new logger with a custom output format
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	// Optional: you can set it as the global logger if you want
	log.Logger = logger

	return logger
}
