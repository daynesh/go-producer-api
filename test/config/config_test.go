package config_test

import (
    "testing"

    "github.com/daynesh/go-producer-api/src/config"
    "github.com/stretchr/testify/assert"
)

func TestConfigInstantiation(t *testing.T) {
    var config = &config.Manager{}

    assert.Equal(t, 0, len(config.BrokerAddresses))
}