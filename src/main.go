package main

import (
	"fmt"
	"os"

	"github.com/daynesh/go-producer-api/src/controllers"
	"github.com/daynesh/go-producer-api/src/utils"
	"github.com/docopt/docopt-go"
	log "github.com/sirupsen/logrus"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	// Load config values
	var config = &utils.Config{}
	err := config.Load(docopt.Parse)
	if err != nil {
		// @todo Move to using a real logger with support for log levels
		fmt.Printf("Error loading configs: %s\n", err)
		os.Exit(1)
	}

	// Now initialize logger
	err = utils.InitLogger(config)
	if err != nil {
		fmt.Printf("Error initializing logger: %s\n", err)
		os.Exit(1)
	}

	// First real log output
	log.Info("Starting " + utils.ApplicationTitle)

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
		log.Fatal("Error executing router")
	}
}
