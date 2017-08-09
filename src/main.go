package main

import (
	"fmt"
	"os"

	"github.com/daynesh/go-producer-api/src/config"
	"github.com/daynesh/go-producer-api/src/controllers"
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/docopt/docopt-go"
)

func main() {
	// Load config values
	var config = &config.Manager{}
	err := config.Load(docopt.Parse)
	if err != nil {
		// @todo Move to using a real logger with support for log levels
		fmt.Printf("Error loading configs: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("Starting go-producer-api")

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
