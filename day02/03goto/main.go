package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto breakTag
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}
	return

breakTag:
	fmt.Println("结束for循环")

	// var flag = false
	// for i := 0; i < 10; i++ {
	// 	for j := 'A'; j < 'Z'; j++ {
	// 		if j == 'C' {
	// 			flag = true
	// 			break
	// 		}
	// 		fmt.Printf("%v-%c\n", i, j)

	// 	}
	// 	if flag == true {
	// 		break
	// 	}
	// }
	// fmt.Println("结束for循环")
}
