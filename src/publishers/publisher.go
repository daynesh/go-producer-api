package publishers

import (
	"encoding/json"
	"fmt"

	"github.com/daynesh/go-producer-api/src/config"
	"github.com/daynesh/go-producer-api/src/requests"
	"gopkg.in/Shopify/sarama.v1"
)

// Publisher provides a means of publishing data to a topic
type Publisher struct {
	producer sarama.AsyncProducer
}

// NewPublisher instantiate a new Publisher instance
func NewPublisher(config *config.Manager) (*Publisher, error) {
	// Try instantiating a producer
	producer, err := sarama.NewAsyncProducer(config.BrokerAddresses, sarama.NewConfig())
	if err != nil {
		return nil, err
	}

	// Now instantiate a Publisher instance
	publisher := &Publisher{producer}

	// Listen for errors and handle accordingly
	go publisher.handleAsyncErrors()

	return publisher, nil
}

// Handle publish errors from producer's error output channel
func (p *Publisher) handleAsyncErrors() {
	for err := range p.producer.Errors() {
		fmt.Println("Publish error detected: ", err)
	}
}

// Publish message asynchronously
func (p *Publisher) Publish(msg requests.PublisherRequest) error {
	// Serialize the message first
	serializedInput, err := json.Marshal(msg)
	if err == nil {
		message := &sarama.ProducerMessage{
			Topic: msg.Header.Topic,
			Value: sarama.ByteEncoder(serializedInput),
		}

		// Now send message to producer input channel
		// TODO: How exactly do we handle errors?
		p.producer.Input() <- message
		return nil
	}

	return err
}
