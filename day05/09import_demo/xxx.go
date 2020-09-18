package main

import (
	"fmt"

	calc "github.com/Moonium/studygo/day05/08calc" // 别名 "路径"
)

var x = 100

const pi = 3.14

func init() {
	fmt.Println("自动执行！")
	fmt.Println(x, pi)
}

func main() {
	ret := calc.Add(18, 20)
	fmt.Println(ret)
}
