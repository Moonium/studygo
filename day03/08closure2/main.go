package main

import "fmt"

// 闭包是一个函数，这个函数包含了它外部作用域的一个变量

func adder1() func(int) int {
	var x = 100
	return func(y int) int {
		x += y
		return x
	}
}

func adder2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	ret1 := adder1()
	fmt.Println(ret1(200))
	ret2 := adder2(100)
	fmt.Println(ret2(200))
}
