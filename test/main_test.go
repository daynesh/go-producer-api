package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/daynesh/go-producer-api/src/controllers"
	"github.com/stretchr/testify/assert"
	"gopkg.in/gin-gonic/gin.v1"
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
