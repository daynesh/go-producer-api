package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Manager contains all config params
type Manager struct {
	BrokerAddresses []string
}

// Load configuration values
func (c *Manager) Load() error {
	// First, set default values
	c.BrokerAddresses = []string{"localhost:9092"}

	// Now override default values with any env values
	return envconfig.Process("producerapi", c)
}
