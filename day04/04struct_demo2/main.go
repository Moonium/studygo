package main

import "fmt"

// 结构体是值类型

type person struct {
	name, gender string
	// age          int
}

func f(x person) {
	x.gender = "女" // 修改的是副本的gender
}

func f2(x *person) {
	// (*x).gender = "女"
	x.gender = "女" // 语法糖，自动根据指针找相应的变量
}

func main() {
	var p person
	p.name = "张三"
	p.gender = "男"
	f(p)
	fmt.Println(p.gender)
	f2(&p)
	fmt.Println(p.gender)

	var p2 = new(person)
	p2.name = "李四"
	(*p2).gender = "保密"
	fmt.Printf("%T\n", p2)
	fmt.Println(p2)
	fmt.Printf("%p\n", p2)

	// key-value初始化
	var p3 = &person{
		name:   "王五",
		gender: "男",
	}
	fmt.Printf("%#v\n", p3)

	// 值列表初始化。值的顺序要和结构体定义时的字段一致
	var p4 = person{
		"小王子",
		"男",
	}
	fmt.Printf("%#v\n", p4)
}
