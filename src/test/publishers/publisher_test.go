package publishers

import (
	"testing"

	"app/publishers"
	"app/utils"

	"github.com/stretchr/testify/assert"
)

func EmptyOptionsParser(string, []string, bool, string, bool, ...bool) (map[string]interface{}, error) {
	return map[string]interface{}{}, nil
}

func TestNewPublisherWithNoBrokers(t *testing.T) {
	var err error
	var config = &utils.Config{}
	err = config.Load(EmptyOptionsParser)
	assert.Equal(t, err, nil)

	// Actual test
	_, err = publishers.NewPublisher(config)

	// Now verify results
	assert.Equal(t, err.Error(), "kafka: client has run out of available brokers to talk to (Is your cluster reachable?)")
}
