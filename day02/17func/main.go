package main

import "fmt"

func sum(x int, y int) (ret int) {
	return x + y
}

func f1(x int, y int) {
	fmt.Println(x + y)
}

func f2() {
	fmt.Println("没有参数没有返回值")
}

func f3() int {
	ret := 3
	return ret
}

func f4(x int, y int) (ret int) {
	ret = x + y // 命名的返回值就相当于在函数中声明了一个变量
	return
}

func f5() (int, string) {
	return 1, "你好"
}

// 连续N个参数类型一致时可以将N-1个简写
func f6(x, y int) (ret int) {
	return x + y
}

// 可变长参数
func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y) // y的类型是切片
}

func main() {
	r := sum(10, 20)
	fmt.Println(r)

	_, n := f5()
	fmt.Println(n)

	f7("张三", 1, 2, 3, 4, 5, 6, 7, 8)
}
