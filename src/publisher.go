package main

import (
    "encoding/json"

    "gopkg.in/Shopify/sarama.v1"
)

type Publisher struct {
    producer sarama.AsyncProducer
}

func NewPublisher() (*Publisher, error) {
    config := sarama.NewConfig()
    producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
    if err != nil {
        return nil, err
    }

    return &Publisher{producer}, nil
}

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