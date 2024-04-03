package main

import (
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	sayhello.sayhello()
	c.JSON(200, gin.H{
		"message": "info",
	})
}

func main() {
	r := gin.Default()
	r.GET("/ping", ping)
	r.Run()
}
