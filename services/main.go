package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	content := os.Getenv("CONTENT")
	port := os.Getenv("PORT")

	server := gin.Default()
	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": content,
		})
	})
	server.Run(port)
}
