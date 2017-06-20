package main

import (
    "fmt"
    "encoding/json"

    "gopkg.in/Shopify/sarama.v1"
)

type Publisher struct {
    producer sarama.AsyncProducer
}

// Instantiate a new Publisher instance
func newPublisher() (*Publisher, error) {
    config := sarama.NewConfig()

    // Try instantiating a producer
    producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
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
func (p *Publisher) publish(msg PublisherPayload) error {
    // Serialize the message first
    serializedInput, err := json.Marshal(msg)
    if err == nil {
        message := &sarama.ProducerMessage{
            Topic: msg.Header.Type,
            Value: sarama.ByteEncoder(serializedInput),
        }

        // Now send message to producer input channel
        // TODO: How exactly do we handle errors?
        p.producer.Input() <- message
        return nil
    } else {
        return err
    }
}