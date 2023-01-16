package main_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	a "gitlab.com/fish895623/golangr"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := a.SetUpRouter()
	gin.SetMode(gin.TestMode)

	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)

	fmt.Print(string(responseData))

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
