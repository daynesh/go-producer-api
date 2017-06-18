package main

import (
    "fmt"

    "gopkg.in/gin-gonic/gin.v1"
)

var publisher *Publisher

// Respond to ping requests
func ping(c *gin.Context) {
    c.JSON(200, gin.H{"message": "pong"})
}

// Publish message from request body
func publishMessage(c *gin.Context) {
    // Gather input params
    var input PublisherPayload
    if c.BindJSON(&input) == nil {

        // TODOs:
        //  - Need to add support for thread safety!
        //  - Need to close connections when done with publisher
        var err error
        if publisher == nil {
            publisher, err = NewPublisher()
            if err != nil {
                c.JSON(500, gin.H{"code": 500, "message": "Internal Server Error", "description": err.Error()})
            }
        }

        // Now publish message
        err = publisher.publish(input)
        if err == nil {
            c.JSON(200, gin.H{
                "status": "OK",
            })
        } else {
            fmt.Println("error: ", err)
        }
    }
}