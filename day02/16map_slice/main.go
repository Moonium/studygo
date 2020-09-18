package main

import "fmt"

func main() {
	// 元素类型为map的切片
	var s1 = make([]map[int]string, 10, 10)
	// 需要对内部map初始化
	s1[0] = make(map[int]string, 1)
	s1[0][100] = "A"
	fmt.Println(s1)

	// 值为切片类型的map
	var m1 = make(map[string][]int, 10)
	m1["张三"] = []int{18, 22, 25}
	fmt.Println(m1)
}
