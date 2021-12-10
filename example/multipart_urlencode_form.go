package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// form表单参数
func main() {
	router := gin.Default()

	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	// 默认端口即为8080
	router.Run(":1010")
}
