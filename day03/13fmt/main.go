package main

import "fmt"

func main() {
	fmt.Print("你好")
	fmt.Print("张三")
	fmt.Println("----------")
	fmt.Println("你好")
	fmt.Println("张三")

	var m1 = make(map[string]int, 1)
	m1["张三"] = 100
	fmt.Printf("%v\n", m1)
	fmt.Printf("%#v\n", m1)

	printPercent(90)
	// 整数->字符
	fmt.Printf("%q\n", 65)
	// 浮点数和复数
	fmt.Printf("%b\n", 3.1415926)
	// 字符串
	fmt.Printf("%q\n", "法外狂徒张三")

	fmt.Printf("%5.2s\n", "张三")

	// 获取用户输入
	// var s string
	// fmt.Scan(&s)
	// fmt.Println("用户输入的内容是", s)

	var (
		name  string
		age   int
		class string
	)
	// fmt.Scanf("%s %d %s\n", &name, &age, &class)
	fmt.Scan(&name, &age, &class)
	fmt.Println(name, age, class)
}

func printPercent(num int) {
	fmt.Printf("%d%%\n", num)
}
