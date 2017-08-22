package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// ServicesJSONFormatter is a custom formatter based on JSONFormatter
type ServicesJSONFormatter struct {
	baseFormatter *log.JSONFormatter
}

// Format provides customized formatting for all log messages
func (sjf *ServicesJSONFormatter) Format(entry *log.Entry) ([]byte, error) {
	// Add additional fields to all log messages
	entry.Data["svc"] = ApplicationTitle

	return sjf.baseFormatter.Format(entry)
}

// InitLogger initializes and wires up how Logger should behave
// throughout the application
func InitLogger(config *Config) error {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&ServicesJSONFormatter{&log.JSONFormatter{}})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log a certain level of severity or above.
	logLevel, err := log.ParseLevel(config.LogLevel)
	if err != nil {
		return err
	}
	log.SetLevel(logLevel)

	return nil
}
