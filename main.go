package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	return r
}

type D struct {
	Load bool   `json:"load"`
	Data string `json:"data"`
}

func aa() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "hello %s", "hello")
	}
}

func Hello() gin.HandlerFunc {
	return func(c *gin.Context) {
		body := D{}
		c.BindJSON(&body)
	}
}

func main() {
	gin.SetMode(gin.DebugMode)

	r := SetUpRouter()

	r.GET("/", aa())
	r.POST("/a", Hello())

	r.Run(":8081")
}
