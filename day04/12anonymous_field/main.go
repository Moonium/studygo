package main

import "fmt"

// 匿名字段，用于字段比较少比较简单的场景

type person struct {
	string
	int
}

func main() {
	p1 := person{
		"张三",
		90,
	}
	fmt.Println(p1)
	fmt.Println(p1.string)
	fmt.Println(p1.int)
}
