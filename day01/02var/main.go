package main

import (
	"fmt"
)

// var name string
// var age int
// var isOk bool

var (
	name string
	age  int
	isOk bool
)

func main() {
	name = "张三"
	age = 16
	isOk = true

	fmt.Print(isOk)
	fmt.Println()
	fmt.Printf("name:%s\n", name)
	fmt.Println(age)

	var s1 string = "王五"
	fmt.Println(s1)

	var s2 = 20
	fmt.Println(s2)

	s3 := "哈哈哈"
	fmt.Println(s3)
}
