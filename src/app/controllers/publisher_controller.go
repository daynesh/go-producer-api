package controllers

import (
	"net/http"

	"app/publishers"
	"app/requests"
	"app/utils"

	"github.com/gin-gonic/gin/binding"
	"gopkg.in/gin-gonic/gin.v1"
)

// PublishController includes config params, a publisher and this
// controllers associated handlers
type PublishController struct {
	Config    *utils.Config
	Publisher *publishers.Publisher
}

// GetPublishController returns a new PublishController
func GetPublishController(config *utils.Config) *PublishController {
	return &PublishController{Config: config}
}

// PublishMessage publishes a message specified in the request body
func (pc *PublishController) PublishMessage(c *gin.Context) {
	// Default response info to return
	responseCode := http.StatusInternalServerError
	var jsonResponse gin.H

	// Gather input params
	var input requests.PublisherRequest
	var err error
	if err = binding.JSON.Bind(c.Request, &input); err == nil {

		// TODOs:
		//  - Need to add support for thread safety!
		//  - Need to close connections when done with publisher
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
			jsonResponse = gin.H{
				"code":        http.StatusInternalServerError,
				"message":     "Internal Server Error",
				"description": err.Error(),
			}
		} else {
			responseCode = http.StatusOK
			jsonResponse = gin.H{"status": "OK"}
		}
	} else {
		responseCode = http.StatusBadRequest
		jsonResponse = gin.H{
			"code":        http.StatusBadRequest,
			"message":     "Bad Request",
			"description": err.Error(),
		}
	}

	c.JSON(responseCode, jsonResponse)
}
