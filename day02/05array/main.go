package main

import (
	"fmt"
)

func main() {
	var a1 [3]bool
	var a2 [4]bool // 长度是数组类型的一部分，不同长度的数组是不同类型

	fmt.Printf("a1:%T a2:%T\n", a1, a2)
	fmt.Println(a1, a2)
	// 具体
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	// 推断
	a10 := [...]int{1, 2, 3, 4, 5, 6, 6, 9}
	fmt.Println(a10)
	// 索引
	a3 := [5]int{0: 1, 4: 2}
	fmt.Println(a3)

	cities := [...]string{"北京", "上海", "深圳"}
	// 根据索引遍历
	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}
	for i, v := range cities {
		fmt.Println(i, v)
	}

	var a11 [3][2]int
	a11 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a11)

	// 多维数组的遍历
	for _, v1 := range a11 {
		// fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	// 数组是值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1)
	fmt.Println(b2)
}
