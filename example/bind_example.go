package main

import "github.com/gin-gonic/gin"

type User struct {
	Name string `json:"name,omitempty" xml:"name"`
	Age  int    `json:"age" xml:"age"`
	Sex  bool   `json:"sex" xml:"sex"`
}

func main() {
	engine := gin.Default()

	engine.Run()
}
