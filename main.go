package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/fish895623/golangr/filemanage"
	ma "gitlab.com/fish895623/golangr/tes"

	"gorm.io/gorm"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Static("/static", "./static")
	router.StaticFile("/favicon.ico", "./favicon.ico")

	return router
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
			url.Data = filemanage.Readfile()
		} else {
			filemanage.Writefile(url.Data)
		}
		c.JSON(http.StatusOK, &url)
	}
}

func main() {
	var db *gorm.DB
	db = ma.InitDb(db)

	router := SetUpRouter()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello %s", "hello")
	})

	router.Run(":8081")
}
