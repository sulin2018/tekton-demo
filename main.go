package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sum(a, b int) int {
	return a + b
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong, hello world! ZA!",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
