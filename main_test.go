package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	rPath := "/user"

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET(rPath, func(c *gin.Context) {
		c.String(http.StatusOK, "hello %s", "hello")
	})
	req, _ := http.NewRequest("GET", rPath, nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "response status should be 200")
	assert.Equal(t, w.Body.String(), "hello hello")
}
