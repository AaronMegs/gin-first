package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	// 增删改查请求的使用
	// GET
	r.GET("/path/:id", func(c *gin.Context) {
		// 取参数占位符值
		id := c.Param("id")
		// 取请求参数值
		//user := c.Query("user")
		// 设置请求参数默认值
		user := c.DefaultQuery("user", "AaronMegs")
		pwd := c.Query("pwd")

		// 返回参数
		c.JSON(http.StatusOK, gin.H{
			"id":   id,
			"user": user,
			"pwd":  pwd,
		})
	})

	// POST请求
	// Form 表单格式
	r.POST("/path", func(c *gin.Context) {
		user := c.DefaultPostForm("user", "AaronMegs")
		pwd := c.PostForm("pwd")

		c.JSON(http.StatusOK, gin.H{
			"user": user,
			"pwd":  pwd,
		})

	})

	// PUT请求 - Form
	r.PUT("path", func(c *gin.Context) {
		user := c.DefaultPostForm("user", "AaronMegs")
		pwd := c.PostForm("pwd")

		c.JSON(http.StatusOK, gin.H{
			"user": user,
			"pwd":  pwd,
		})
	})

	// DELETE请求
	r.DELETE("path/:id", func(c *gin.Context) {
		id := c.Param("id")

		c.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	})

	//r.Run() // 监听并在localhost:8080 上启动服务
	// 修改端口号
	r.Run(":1010")

	// net/http
	//http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {})
	//http.HandleFunc("/hello", testHttp)
	//log.Fatal(http.ListenAndServe("localhost:8080", nil))

}

func testHttp(w http.ResponseWriter, req *http.Request) {
	// 获取url中的参数
	//number := req.URL.Query().Get("number")
	fmt.Fprintf(w, "hello world")
}
