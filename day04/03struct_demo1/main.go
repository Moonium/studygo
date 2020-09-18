package main

import "fmt"

type person struct {
	name   string
	age    int
	gender bool
	hobby  []string
}

func main() {
	var p person
	p.name = "张三"
	p.age = 18
	p.gender = true
	p.hobby = []string{"篮球", "足球", "双色球"}
	fmt.Println(p)
	fmt.Printf("%T\n", p)
	fmt.Println(p.name)

	var p2 person
	p2.name = "李四"
	p2.age = 22
	fmt.Printf("type:%T value:%v\n", p2, p2)
	fmt.Println(p2.hobby == nil)

	// 匿名结构体：多用于一些临时场景
	var s struct {
		x string
		y int
	}
	s.x = "嘿嘿嘿"
	s.y = 100
	fmt.Printf("type:%T value:%v\n", s, s)
}
