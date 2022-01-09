package main

import "fmt"

const (
	// iota (读作 约塔 希腊字母）从0开始计数，依次正整数
	_ = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

const (
	// iota (读作 约塔 希腊字母）从0开始计数，依次负整数
	_ = -iota
	Hello
	World
)

func main() {
	var week interface{}
	week = Monday
	// 1
	fmt.Println(week)
	fmt.Printf("%v", week)
	fmt.Println()
	// -1
	fmt.Println(Hello)
}
