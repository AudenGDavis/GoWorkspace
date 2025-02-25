package main

import (
	"github.com/gin-gonic/gin"
)

func basicGet(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello, World!", "also": "hi there"})
}

func main() {
	r := gin.Default()
	r.GET("/hello", basicGet)

	r.Run(":8080")
}
