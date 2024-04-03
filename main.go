package main

import (
	"golang/sayhello"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	sayhello.Sayhello()
	c.JSON(200, gin.H{
		"message": "info",
	})
}

func main() {
	r := gin.Default()
	r.GET("/ping", ping)
	r.Run()
}
