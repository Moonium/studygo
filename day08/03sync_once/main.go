package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var once sync.Once

func f1(ch1 chan<- int) { // 单向通道
	defer wg.Done()
	for i := 1; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func f2(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	f := func() { close(ch2) } // 用闭包封装
	once.Do(f)                 // 确保某一操作只执行一次
}

func main() {
	a := make(chan int, 100)
	b := make(chan int, 100)
	wg.Add(3)
	go f1(a)
	go f2(a, b)
	go f2(a, b)
	wg.Wait()
	for ret := range b {
		fmt.Println(ret)
	}

}