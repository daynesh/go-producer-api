package main

import (
	"fmt"
	"os"

	"github.com/daynesh/go-producer-api/src/config"
	"github.com/daynesh/go-producer-api/src/controllers"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	fmt.Println("Starting go-producer-api")

	// Load config values
	var config = &config.Manager{}
	err := config.Load()
	if err != nil {
		// @todo Move to using a real logger with support for log levels
		fmt.Println("Error loading configs")
		os.Exit(1)
	}

	// Instantiate an Engine instance
	router := gin.Default()

	// Instantiate controllers
	var publishController = controllers.GetPublishController(config)
	var pingController = controllers.GetPingController()

	// Route definitions
	router.GET("/ping", pingController.Ping)
	router.POST("/publish", publishController.PublishMessage)

	err = router.Run()
	if err != nil {
		fmt.Println("Error executing router")
	}
}
