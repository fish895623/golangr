package ma_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	ma "github.com/fish895623/golangr/tes"
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
	router, req, _ := DefaultGetSetup("/user", ma.Aa())

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "response status should be 200")
	assert.Equal(t, w.Body.String(), "hello hello")
}

func TestFF(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/", ma.Hello())
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer([]byte(`{"load": true, "data": "asdaaaa"}`)))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, w.Body.String(), `{"load":false,"data":"Hello"}`)
}
func TestFalse(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/", ma.Hello())
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer([]byte(`{"load": false, "data": "asdaaaa"}`)))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, w.Body.String(), `{"load":true,"data":"Hello"}`)
}
