package main

import (
	"fmt"

	"github.com/daynesh/go-producer-api/src/config"
	"github.com/daynesh/go-producer-api/src/controllers"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	fmt.Println("Starting go-producer-api")

	// Load config values
	var config = &config.Manager{}
	config.Load()

	// Instantiate an Engine instance
	router := gin.Default()

	// Instantiate controllers
	var publishController = controllers.GetPublishController(config)
	var pingController = controllers.GetPingController()

	// Route definitions
	router.GET("/ping", pingController.Ping)
	router.POST("/publish", publishController.PublishMessage)

	router.Run()
}
