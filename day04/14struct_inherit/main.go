package main

import "fmt"

// 结构体模拟实现其他语言中的“继承”

type animal struct {
	name string
}

// 给animal实现一个“动”的方法
func (a animal) move() {
	fmt.Printf("%s在动\n", a.name)
}

type dog struct {
	feet int8
	animal
}

// 给dog实现一个“汪汪汪”的方法
func (d dog) wang() {
	fmt.Printf("%s在汪汪汪！", d.name)
}

func main() {
	d1 := dog{
		4,
		animal{
			"旺财",
		},
	}
	fmt.Println(d1)
	d1.move()
	d1.wang()
}
