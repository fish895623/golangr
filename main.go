package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/fish895623/golangr/tes"
	"gorm.io/gorm"
)

var router *gin.Engine

func setupConfig() {
	router = gin.Default()
	router.SetTrustedProxies([]string{"192.168.0.2"})

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")
	router.StaticFile("/favicon.ico", "./favicon.ico")
}

func main() {
	var db *gorm.DB
	db = ma.InitDb(db)

	setupConfig()

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
		var url struct {
			Name string `json:"data"`
		}
		c.BindJSON(&url)
		c.JSON(http.StatusOK, &url)
		fmt.Print(url.Name)
	})

	router.Run()
}
