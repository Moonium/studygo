package main

import "fmt"

var x = 100

func f1() {
	// 函数中查找变量的顺序
	// 1.现在函数内查找
	// 2.找不到就在函数的外部查找，一直找到全局
	name := "张三"
	fmt.Println(x, name)
}

func main() {
	f1()
}
