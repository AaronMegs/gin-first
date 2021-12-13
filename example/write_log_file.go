package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

// 将日志写到文件中
func main() {
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file
	file, err := os.Create("gin.log")
	if err != nil {
		fmt.Printf("创建文件失败：%s ", err)
		//errors.New("创建文件失败")
		return
	}
	// Only write the logs to file
	//gin.DefaultWriter = io.MultiWriter(file)

	// Use the following code if you need to write the logs to file and console at the same time.
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":8080")
}
