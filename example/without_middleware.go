package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Blank Gin  without middleware by default
	//r := gin.New()

	// Default with the Logger and Recovery middleware already attached
	r := gin.Default()

	r.GET("/hello/:name", func(c *gin.Context) {
		param := c.Param("name")

		c.JSON(http.StatusOK, gin.H{
			"name": param,
		})

	})

	r.Run(":8080")
}
