package main

import "fmt"

type myInt int

// 给自定义类型加方法
// 不能给别的包里的类型添加方法
func (m myInt) sayHello() {
	fmt.Println("我是一个int")
}

func main() {
	m := myInt(100)
	m.sayHello()
}
