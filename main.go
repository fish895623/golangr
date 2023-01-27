package main

import (
	ma "github.com/fish895623/golangr/tes"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)

	r := gin.Default()

	r.GET("/", ma.Aa())
	r.POST("/a", ma.Hello())

	r.Run(":8081")
}
