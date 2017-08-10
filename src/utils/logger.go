package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// Logger will be used for all logging of messages
var Logger *log.Entry

// InitLogger initializes and wires up how Logger should behave
// throughout the application
func InitLogger(config *Config) error {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log a certain level of severity or above.
	logLevel, err := log.ParseLevel(config.LogLevel)
	if err != nil {
		return err
	}
	log.SetLevel(logLevel)

	// Add custom fields to all log messages
	Logger = log.WithField("svc", ApplicationTitle)

	return nil
}
