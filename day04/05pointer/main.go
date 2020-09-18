package main

import "fmt"

func main() {
	var a int
	a = 100
	b := &a
	fmt.Printf("a:%T b:%T\n", a, b)
	// 将a的十六进制内存地址打印出来
	fmt.Printf("%p\n", b)
	fmt.Printf("%v\n", b)
	// 将b的十六进制内存地址打印出来
	fmt.Printf("%p\n", &b)
}
