package utils

import (
	"errors"
	"os"
	"path"
	"strings"

	"github.com/kelseyhightower/envconfig"
	"github.com/mitchellh/mapstructure"
)

// Config contains all config params
type Config struct {
	BrokerAddresses []string `envconfig:"BROKER_ADDRESSES" default:"localhost:9092"`
	LogLevel        string   `envconfig:"LOG_LEVEL" default:"info" mapstructure:"--log-level"`
}

// Constant definitions
const (
	ApplicationTitle = "Producer API"
	EnvConfigPrefix  = "PRODUCERAPI"
)

// OptionsParse type is used for dependency injecting the docopt.Parse() logic
type OptionsParse func(string, []string, bool, string, bool, ...bool) (map[string]interface{}, error)

// Usage information for specifying config data
func getUsage() string {
	me := path.Base(os.Args[0])

	return ApplicationTitle + `
Usage:
	` + me + ` -h | --help
	` + me + ` [--broker-addresses=<addr1,addr2>] [--log-level=<debug|info|warning|error|fatal|panic>]

Options:
	-h --help   Show this screen.
	--broker-addresses=<addr>
	--log-level=<loglevel>`
}

// Decode key from optionsMap and write to destination
func decodeArrayOfStrings(optionsMap map[string]interface{}, key string, destination *[]string) error {
	// First check if key exists...if not, nothing to do
	value, exists := optionsMap[key]
	if exists && value != nil {
		// Type assert to string
		input, found := optionsMap[key].(string)
		if !found {
			return errors.New("Specified option " + key + " could not be converted to a string")
		}

		// Now split string into an array of strings
		arrayOfStrings := strings.Split(input, ",")

		// Finally, store the array into the destination
		*destination = arrayOfStrings
	}

	return nil
}

// Load configuration values
func (c *Config) Load(optionsParse OptionsParse) error {
	// Now override default values with any env values
	err := envconfig.Process(EnvConfigPrefix, c)
	if err != nil {
		return err
	}

	// Finally, attempt to override with any specified CLI options
	arguments, err := optionsParse(getUsage(), nil, true, ApplicationTitle, false)
	if err != nil {
		return err
	}
	err = mapstructure.Decode(arguments, c)
	if err != nil {
		return err
	}

	// For some reason, arrays of strings are not decoding correctly
	// See here: https://github.com/mitchellh/mapstructure/issues/91
	// So instead we must decode string arrays manually here
	err = decodeArrayOfStrings(arguments, "--broker-addresses", &c.BrokerAddresses)

	return err
}
