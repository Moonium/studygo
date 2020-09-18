package main

import "fmt"

// 二进制实际用处

const (
	eat   int = 4
	sleep int = 2
	beat  int = 1
)

// 左边的1表示吃饭
// 中间的1表示睡觉
// 右边的1表示打豆豆
func f(arg int) {
	fmt.Printf("%b\n", arg)
}

func main() {
	f(eat | beat)
	f(eat | beat | sleep)
}
