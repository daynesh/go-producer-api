package config

import (
	"github.com/kelseyhightower/envconfig"
)

type ConfigManager struct {
	BrokerAddresses []string
}

// Load configuration values
func (c *ConfigManager) Load() error {
	// First, set default values
	c.BrokerAddresses = []string{"localhost:9092"}

	// Now override default values with any env values
	return envconfig.Process("producerapi", c)
}
