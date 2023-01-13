package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	ma "gitlab.com/fish895623/golangr/tes"
	"gorm.io/gorm"
)

func setupConfig(router *gin.Engine) *gin.Engine {
	router.SetTrustedProxies([]string{"192.168.0.2"})
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")
	router.StaticFile("/favicon.ico", "./favicon.ico")

	return router
}

func hello() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"header.html",
			gin.H{
				"title": "Home Page",
			},
		)
	}
}

type D struct {
	Load bool   `json:"load"`
	Data string `json:"data"`
}

func api() gin.HandlerFunc {
	return func(c *gin.Context) {
		url := D{}
		c.BindJSON(&url)
		if url.Load {
			url.Data = readfile()
		} else {
			writefile(url.Data)
		}
		c.JSON(http.StatusOK, &url)
	}
}

func readfile() string {
	dat, _ := os.ReadFile("./write")
	return string(dat)
}

func writefile(a string) {
	data := []byte(a)
	os.WriteFile("./write", data, 0644)
}

func main() {
	var db *gorm.DB
	db = ma.InitDb(db)

	var router *gin.Engine = gin.Default()
	router = setupConfig(router)

	a := router.Group("/a")
	a.GET("/", hello())
	router.POST("/api", api())

	router.Run()
}
