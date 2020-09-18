package main

import "fmt"

func f1() {
	fmt.Println("Hello, 张三")
}

func f2(name string) {
	fmt.Printf("Hello, %s\n", name)
}

func f3(x int, y int) int {
	sum := x + y
	return sum
}

func f4(x, y int) int {
	return x + y
}

func f5(title string, y ...int) int {
	fmt.Println(y)
	return 1
}

func f6(x, y int) (sum int) {
	sum = x + y // 使用命名的返回值，函数中可以直接使用返回值变量
	return      // 使用命名的返回值，return后可以省略返回值变量
}

func f7(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

func main() {
	f1()
	f1()
	f2("李四")
	ret := f3(100, 200)
	fmt.Println(ret)
	f5("张三", 1, 2, 3, 4, 5, 6, 7, 8, 9)
}
