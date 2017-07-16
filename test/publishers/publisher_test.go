package publisher_test

import (
    "testing"

    "github.com/daynesh/go-producer-api/src/config"
    "github.com/daynesh/go-producer-api/src/publishers"
    "github.com/stretchr/testify/assert"
)

func TestNewPublisher(t *testing.T) {
    var publisher *publishers.Publisher
    var err error
    var config = &config.Manager{}

    publisher, err = publishers.NewPublisher(config)

    assert.Equal(t, publisher, nil)
    assert.Equal(t, err, nil)
}