package publisher_test

import (
	"testing"

	"github.com/daynesh/go-producer-api/src/config"
	"github.com/daynesh/go-producer-api/src/publishers"
	"github.com/stretchr/testify/assert"
)

func EmptyOptionsParser(string, []string, bool, string, bool, ...bool) (map[string]interface{}, error) {
	return map[string]interface{}{}, nil
}

func TestNewPublisherWithNoBrokers(t *testing.T) {
	var err error
	var config = &config.Manager{}
	config.Load(EmptyOptionsParser)

	// Actual test
	_, err = publishers.NewPublisher(config)

	// Now verify results
	assert.Equal(t, err.Error(), "kafka: client has run out of available brokers to talk to (Is your cluster reachable?)")
}
