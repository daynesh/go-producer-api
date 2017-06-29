package controllers

import (
	"gopkg.in/gin-gonic/gin.v1"
)

// PingController provides a means of referencing handlers
// grouped by this specific controller
type PingController struct{}

// GetPingController returns a new PingController
func GetPingController() *PingController {
	return &PingController{}
}

// Ping handles a ping request by returning a pong
func (pc *PingController) Ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}
