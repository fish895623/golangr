package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func DefaultGetSetup(rPath string, a gin.HandlerFunc) (*gin.Engine, *http.Request, error) {

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET(rPath, a)
	req, err := http.NewRequest("GET", rPath, nil)

	return router, req, err
}

func TestPingRoute(t *testing.T) {
	router, req, _ := DefaultGetSetup("/user",
		func(c *gin.Context) {
			c.String(http.StatusOK, "hello %s", "hello")
		},
	)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "response status should be 200")
	assert.Equal(t, w.Body.String(), "hello hello")
}
