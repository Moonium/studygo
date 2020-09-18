package main

import (
	"fmt"
	"sync"
)

var a []int
var b chan int // 需要指定通道中元素的类型
var wg sync.WaitGroup

func noBufChannel() {
	fmt.Println(b)
	b = make(chan int) // 通道的初始化
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("后台goroutine从通道b中取到了", x)
	}()
	b <- 10 // 死锁
	fmt.Println("数值发送到通道b中了")
	wg.Wait()
}

func bufChannel() {
	fmt.Println(b)
	b = make(chan int, 1) // 通道的初始化
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	x := <-b
	// 	fmt.Println("后台goroutine从通道b中取到了", x)
	// }()
	b <- 10
	// b <- 20 // 死锁
	fmt.Println("数值发送到通道b中了")
	// wg.Wait()
	x := <-b
	fmt.Println("从通道b中取到了", x)
	close(b)
}

func main() {

	// b = make(chan int, 16) // 带缓冲区的通道初始化
	// fmt.Println(b)
	bufChannel()
}
