package main

import (
	"fmt"
)

func main() {
	// age := 19
	// if age > 18 {
	// 	fmt.Println("澳门首家线上赌场开业啦！")
	// } else {
	// 	fmt.Println("该写暑假作业了！")
	// }

	// if age > 35 {
	// 	fmt.Println("人到中年")
	// } else if age > 18 {
	// 	fmt.Println("青年")
	// } else {
	// 	fmt.Println("好好学习！")
	// }

	// 作用域
	if age := 19; age > 18 {
		fmt.Println("澳门首家线上赌场开业啦！")
	} else {
		fmt.Println("该写暑假作业了！")
	}

	// fmt.Println(age)
}
