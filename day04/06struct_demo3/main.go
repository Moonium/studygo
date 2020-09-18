package main

import "fmt"

// 结构体占据一段连续的内存
type x struct {
	a int8
	b int8
	c string
}

func main() {
	m := x{
		a: int8(10),
		b: int8(20),
		c: "嘿嘿嘿", // 内存对齐
	}
	fmt.Printf("%p\n", &(m.a))
	fmt.Printf("%p\n", &m.b)
	fmt.Printf("%p\n", &m.c)
}
