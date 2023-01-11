package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type Body struct {
	Name string `json:"data"`
}

var router *gin.Engine

func main() {

	router = gin.Default()
	router.SetTrustedProxies([]string{"192.168.0.2"})

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")
	router.StaticFile("/favicon.ico", "./favicon.ico")

	a := router.Group("/a")
	a.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"header.html",
			gin.H{
				"title": "Home Page",
			},
		)
	})
	router.POST("/api", func(c *gin.Context) {
		body := Body{}
		fmt.Println(body)
		c.BindJSON(&body)
		c.JSON(http.StatusAccepted, &body)
	})

	router.Run()
}
