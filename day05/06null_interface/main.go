package main

import "fmt"

// 空接口

// 空接口作为函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

func main() {
	// interface：关键字
	// interface{}：空接口类型
	m1 := make(map[string]interface{}, 16)
	m1["name"] = "张三"
	m1["age"] = 22
	m1["married"] = true
	m1["hobby"] = [...]string{"唱", "跳", "rap"}
	fmt.Println(m1)

	show(nil)
	show(false)
	show(m1)
}
