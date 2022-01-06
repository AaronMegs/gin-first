package main

import "fmt"

type user struct {
	name string
	age  int
}

func main() {
	// 接口 assert 断言 两种方式
	params := make([]interface{}, 3)
	params[0] = 88                // 整型
	params[1] = "hello"           // 字符串
	params[2] = user{"aaron", 25} // 自定义结构体
	// 1、comma-OK 断言
	for index, v := range params {
		if _, ok := v.(int); ok {
			fmt.Printf("Params[%d] is int type.\n", index)
		} else if _, ok := v.(string); ok {
			fmt.Printf("Params[%d] is string type.\n", index)
		} else if _, ok := v.(user); ok {
			fmt.Printf("Params[%d] is custom struct type.\n", index)
		} else {
			fmt.Printf("Params[%d] unknow type.\n", index)
		}
	}

	// 2、switch 断言
	for index, v := range params {
		switch value := v.(type) {
		case int:
			fmt.Printf("Params[%d] is int type, value is %d \n", index, value)
		case string:
			fmt.Printf("Params[%d] is string type, value is %s \n", index, value)
		case user:
			fmt.Printf("Params[%d] is user type, value is %v \n", index, value)
		default:
			fmt.Printf("Params[%d] is unknow.\n", index)
		}
	}

}
