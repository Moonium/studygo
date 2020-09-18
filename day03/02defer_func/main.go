package main

import "fmt"

// defer多用于函数结束之前释放资源（文件句柄、数据库连接、socket连接）
func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("嘿嘿嘿") // defer将它后面的语句延迟到函数即将返回的时候执行
	defer fmt.Println("呵呵呵") // 一个函数中多个defer按照先进后出的顺序进行延迟
	defer fmt.Println("哈哈哈")
	fmt.Println("end")
}

func main() {
	deferDemo()
}
