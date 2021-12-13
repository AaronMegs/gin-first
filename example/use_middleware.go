package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// AuthRequired 自定义中间件
func AuthRequired() gin.HandlerFunc {

	return func(c *gin.Context) {
		// 请求处理之前之前执行
		fmt.Println("请求前执行")

		c.Next()

		// 请求处理之后执行
		fmt.Println("请求后执行")
	}
}

func MyBenchLogger() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

// Using middleware
func main() {
	// Create a router without any middleware by default
	router := gin.Default()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default, gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was onw
	router.Use(gin.Recovery())

	// Per router middleware, you can add as many as you desire
	//router.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	// Authorization group
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	authorized := router.Group("/")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", func(c *gin.Context) {
			fmt.Println("login api")
		})
		authorized.POST("/submit", func(c *gin.Context) {
			fmt.Println("submit api")
		})
		authorized.POST("/read", func(c *gin.Context) {
			fmt.Println("read api")
		})

		// nested group -- 内嵌组
		testing := authorized.Group("/testing")
		testing.GET("/analytics", func(c *gin.Context) {
			fmt.Println("analytics api")
			c.JSON(200, gin.H{
				"sub group": "analytics api",
			})
		})
	}

	// Listen and server on 0.0.0.0:8080
	router.Run(":8080")
}
