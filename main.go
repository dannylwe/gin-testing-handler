package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var PORT string

func init(){
	PORT = "9005"
}

func main() {
	r := setupRouter()
	r.Run(":" + PORT) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	return r
}
