package main

import (
    "fmt"

    "gopkg.in/gin-gonic/gin.v1"
)

func main() {
    fmt.Println("Starting go-producer-api")

    // Load config values
    config.Load()

    // Instantiate an Engine instance
    router := gin.Default()

    // Route definitions
    router.GET("/ping", ping)
    router.POST("/publish", publishMessage)

    router.Run()
}
