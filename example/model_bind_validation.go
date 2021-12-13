package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Model binding and validation
// Gin uses go-playground/validator/v10 for validation. Check the full docs on tags usage here. https://godoc.org/github.com/go-playground/validator#hdr-Baked_In_Validators_and_Tags
// 2 type
// Must bind
// Should bind - usually

// Login Binding from JSON / XML / FORM
type Login struct {
	User     string `form:"user" json:"user" xml:"user" binding:"required"` // binding:"-" 获取不设置 binding:"" tag 时表示该参数为非必填项
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
	//Password string `form:"password" json:"password" xml:"password"`
	//Password string `form:"password" json:"password" xml:"password" binding:"-"`
}

func main() {
	router := gin.Default()

	// Example for binding JSON ({"user": "manu", "password": "123"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "unauthorized",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Example for binding XML (
	//	<?xml version="1.0" encoding="UTF-8"?>
	//	<root>
	//		<user>manu</user>
	//		<password>123</password>
	//	</root>)
	router.POST("/loginXML", func(c *gin.Context) {
		var xml Login
		if err := c.ShouldBindXML(&xml); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if xml.User != "manu" || xml.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are login in"})
	})

	// Example for binding a HTML form (user=manu&password=123)
	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		if err := c.ShouldBind(&form); err != nil {
			//if err := c.ShouldBindQuery(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.User != "manu" || form.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are login in"})
	})

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}
