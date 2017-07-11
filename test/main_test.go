package main_test

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "gopkg.in/gin-gonic/gin.v1"
    "github.com/stretchr/testify/assert"
    "github.com/daynesh/go-producer-api/src/controllers"
)

func TestPing(t *testing.T) {
    testRouter := gin.Default()
    recorder := httptest.NewRecorder()

    // Set up route and handler
    var pingController = controllers.GetPingController()
    testRouter.GET("/ping", pingController.Ping)

    // Send ping request
    r, _ := http.NewRequest("GET", "http://localhost:8080/ping", nil)
    testRouter.ServeHTTP(recorder, r)

    // Verify
    assert.Equal(t, 200, recorder.Code)
    assert.Equal(t, "{\"message\":\"pong\"}", recorder.Body.String())
}