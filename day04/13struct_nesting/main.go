package main

import "fmt"

// 结构体嵌套

type address struct {
	province string
	city     string
}

type workPlace struct {
	province string
	city     string
}

type person struct {
	name string
	age  int
	// province string
	// city     string
	addr address
}

type company struct {
	name string
	// province string
	// city     string
	address // 匿名嵌套
	workPlace
}

func main() {
	p1 := person{
		"张三",
		22,
		address{
			"山东",
			"威海",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.addr.city)

	c1 := company{
		"荣华公司",
		address{
			"甘肃",
			"武威",
		},
		workPlace{
			"甘肃",
			"武威",
		},
	}
	// fmt.Println(c1.city) // 先在自己结构体查找，找不到去匿名嵌套的结构体查找该字段
	fmt.Println(c1.address.city)
	fmt.Println(c1.workPlace.city)
}
