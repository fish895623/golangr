package main

import (
	"net/http"

	"github.com/fish895623/golangr/filemanage"
	ma "github.com/fish895623/golangr/tes"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()

	return r
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

func aa() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "hello %s", "hello")
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	var db *gorm.DB
	db = ma.InitDb(db)

	r := SetUpRouter()

	r.GET("/", aa())

	r.Run(":8081")
}
