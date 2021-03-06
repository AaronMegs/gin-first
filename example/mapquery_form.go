package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// POST /post?ids[a]=1234&ids[b]=hello HTTP/1.1
	//Content-Type: application/x-www-form-urlencoded
	//names[first]=thinkerou&names[second]=tianou
	router.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		c.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})

	})

	router.Run(":8080")
}
