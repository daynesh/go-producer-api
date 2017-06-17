package main

import (
	"fmt"
    "encoding/json"

	"gopkg.in/gin-gonic/gin.v1"
    "gopkg.in/Shopify/sarama.v1"
)

type Header struct {
    Timestamp int64     `json:"timestamp" binding:"required"`
    Type      string    `json:"type" binding:"required"`        // Topic name
    DeviceId  string    `json:"deviceId" binding:"required"`
    UserId    string    `json:"userId" binding:"required"`
}

/* Payload should contain the following structure
    {
        body: {...},                    // raw payload data as JSON
        header: {
            timestamp: 1497406310,      // in Unix epoch time
            type: "sometopic",          // pub-sub topic
            deviceId: "asdf-zxcv-qwer"  // unique identifier for device
            userId: "0123456789"        // if signed in (optional)
        }
    }
*/
type PublisherPayload struct {
    Header  Header           `json:"header" binding:"required"`
    Body    *json.RawMessage `json:"body" binding:"required"`
}

func main() {
	fmt.Println("Starting go-producer-api")

    // Instantiate an Engine instance
	router := gin.Default()
    config := sarama.NewConfig()
    producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
    if err != nil {
        panic(err)
    }

    // Route definitions
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

    router.POST("/publish", func(c *gin.Context) {
        // Gather input params
        var input PublisherPayload
        if c.BindJSON(&input) == nil {
            // publish message
            serializedInput, err := json.Marshal(input)
            if err == nil {
                message := &sarama.ProducerMessage{Topic: input.Header.Type, Value: sarama.ByteEncoder(serializedInput)}
                producer.Input() <- message

                // Now send response
                c.JSON(200, gin.H{"status": "OK"})
            } else {
                fmt.Println("error: ", err)
            }
        }
    })

	router.Run()
}
