package ma

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type D struct {
	Load bool   `json:"load"`
	Data string `json:"data"`
}

func Aa() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "hello %s", "hello")
	}
}

func Hello() gin.HandlerFunc {
	return func(c *gin.Context) {
		body := D{}
		c.BindJSON(&body)
		re := D{}
		re.Data = "Hello"
		if body.Load == false {
			re.Load = true
		}
		c.JSON(http.StatusAccepted, re)
	}
}
