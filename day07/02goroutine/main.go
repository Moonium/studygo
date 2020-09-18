package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Println("hello", i)
}

// 程序启动以后会创建一个main goroutine
func main() {
	for i := 0; i < 100; i++ {
		// go hello(i) // 开启一个单独的goroutine去执行这个任务
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	fmt.Println("main")
	// main函数结束了，main启动的goroutine都会结束
	time.Sleep(time.Second)

}
