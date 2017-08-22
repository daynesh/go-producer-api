package utils

import (
	"testing"

	"app/utils"
	"github.com/stretchr/testify/assert"
)

func EmptyOptionsParser(string, []string, bool, string, bool, ...bool) (map[string]interface{}, error) {
	return map[string]interface{}{}, nil
}

func OptionsParserWithMultipleBrokerAddresses(string, []string, bool, string, bool, ...bool) (map[string]interface{}, error) {
	return map[string]interface{}{
		"--broker-addresses": "localhost:9092,localhost:9093",
	}, nil
}

func TestConfigInstantiation(t *testing.T) {
	var config = &utils.Config{}

	assert.Equal(t, 0, len(config.BrokerAddresses))
}

func TestLoadDefaultConfig(t *testing.T) {
	var config = &utils.Config{}

	err := config.Load(EmptyOptionsParser)

	// Now verify expectations
	assert.Nil(t, err)
	assert.Equal(t, []string{"localhost:9092"}, config.BrokerAddresses, "default broker address check")
}

func TestLoadArrayOfBrokerAddresses(t *testing.T) {
	var config = &utils.Config{}

	err := config.Load(OptionsParserWithMultipleBrokerAddresses)

	// Now verify expectations
	assert.Nil(t, err)
	assert.Equal(t, []string{"localhost:9092", "localhost:9093"}, config.BrokerAddresses, "default broker address check")
}
