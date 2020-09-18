package main

import "fmt"

// 构造函数

type person struct {
	name string
	age  int
}

type dog struct {
	name string
}

// 构造函数拷贝结构体还是结构体指针
// 结构体比较大的时候尽量使用结构体指针减小程序的运行开销
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func newDog(name string) dog {
	return dog{
		name: name,
	}
}

func main() {
	p1 := newPerson("张三", 18)
	p2 := newPerson("李四", 22)
	fmt.Println(*p1, *p2)

	d1 := newDog("旺旺")
	fmt.Println(d1)
}
