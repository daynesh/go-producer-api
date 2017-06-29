package controllers

import (
    "net/http"

    "github.com/daynesh/go-producer-api/src/config"
    "github.com/daynesh/go-producer-api/src/requests"
    "github.com/daynesh/go-producer-api/src/publishers"
    "gopkg.in/gin-gonic/gin.v1"
    "github.com/gin-gonic/gin/binding"
)

type PublishController struct {
    Config *config.ConfigManager
    Publisher *publishers.Publisher
}

func GetPublishController(config *config.ConfigManager) *PublishController {
    return &PublishController{Config: config}
}

// Publish message from request body
func (pc *PublishController) PublishMessage(c *gin.Context) {
    // Default response info to return
    responseCode := http.StatusInternalServerError
    jsonResponse := gin.H{}

    // Gather input params
    var input requests.PublisherRequest
    if err := binding.JSON.Bind(c.Request, &input); err == nil {

        // TODOs:
        //  - Need to add support for thread safety!
        //  - Need to close connections when done with publisher
        var err error
        if pc.Publisher == nil {
            pc.Publisher, err = publishers.NewPublisher(pc.Config)
        }

        if err == nil {
            // Now publish message
            err = pc.Publisher.Publish(input)
        }

        // Handle errors that occurred in either creating a new publisher
        // OR publishing a message
        if err != nil {
            jsonResponse = gin.H{"code": http.StatusInternalServerError,
                "message": "Internal Server Error",
                "description": err.Error(),
            }
        } else {
            responseCode = http.StatusOK
            jsonResponse = gin.H{"status": "OK"}
        }
    } else {
        responseCode = http.StatusBadRequest
        jsonResponse = gin.H{
            "code": http.StatusBadRequest,
            "message": err.Error(),
        }
    }

    c.JSON(responseCode, jsonResponse)
}