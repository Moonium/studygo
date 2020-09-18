package main

import "fmt"

type myInt int

// 1. myInt(100)是个啥
func (m myInt) sayHello() {
	fmt.Println("我是一个int")
}

type person struct {
	name string
	age  int
}

func main() {
	// var x int32 = 10
	x := int32(10)
	fmt.Println(x)

	m := myInt(100)
	m.sayHello()

	// 2. 结构体初始化
	p := person{
		name: "张三",
		age:  22,
	}
	fmt.Println(p)
}

// 3. 为什么要有构造函数
func newPerson(name string, age int) *person {
	return &person{
		name,
		age,
	}
}
