package controllers

import (
    "gopkg.in/gin-gonic/gin.v1"
)

type PingController struct {}

func GetPingController() *PingController {
    return &PingController{}
}

func (pc *PingController) Ping(c *gin.Context) {
    c.JSON(200, gin.H{"message": "pong"})
}