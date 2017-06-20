package main

import (
    "gopkg.in/gin-gonic/gin.v1"
)

var publisher *Publisher

// Respond to ping requests
func ping(c *gin.Context) {
    c.JSON(200, gin.H{"message": "pong"})
}

// Publish message from request body
func publishMessage(c *gin.Context) {
    // Default response info to return
    responseCode := 500
    jsonResponse := gin.H{}

    // Gather input params
    var input PublisherPayload
    if c.BindJSON(&input) == nil {

        // TODOs:
        //  - Need to add support for thread safety!
        //  - Need to close connections when done with publisher
        var err error
        if publisher == nil {
            publisher, err = newPublisher()
        }

        if err == nil {
            // Now publish message
            err = publisher.publish(input)
        }

        // Handle errors that occurred in either creating a new publisher
        // OR publishing a message
        if err != nil {
            jsonResponse = gin.H{"code": 500, "message": "Internal Server Error", "description": err.Error()}
        } else {
            responseCode = 200
            jsonResponse = gin.H{"status": "OK"}
        }
    }

    c.JSON(responseCode, jsonResponse)
}