package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var i = 5
	for ; i <= 10; i++ {
		fmt.Println(i)

	}

	// var i = 5
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// 死循环
	// for {
	// 	fmt.Println(123)
	// }

	s := "Hello 张三"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
	}

	// 哑元变量
	for _, v := range s {
		fmt.Printf("%c\n", v)
	}
}
