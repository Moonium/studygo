package main

import "fmt"

// 引出接口的实例

// 定义一个能“叫”的类型
type speaker interface {
	speak() // 只要实现了speak方法的变量都是speaker类型，方法签名
}

type cat struct{}

type dog struct{}

type person struct{}

func (c cat) speak() {
	fmt.Println("喵喵喵！")
}

func (d dog) speak() {
	fmt.Println("汪汪汪！")
}

func (p person) speak() {
	fmt.Println("啊啊啊！")
}

func da(x speaker) {
	// 传进来一个参数，传进来谁就打谁
	x.speak() // 挨了打就要叫
}

func main() {
	var c1 cat
	var d1 dog
	var p1 person

	da(c1)
	da(d1)
	da(p1)

	var ss speaker // 定义一个接口类型的变量
	ss = c1
	ss = d1
	ss = p1
	fmt.Println(ss)
}
