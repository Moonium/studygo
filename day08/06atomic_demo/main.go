package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	x    int64
	wg   sync.WaitGroup
	lock sync.Mutex
)

func add() {
	// x++
	// lock.Lock()
	atomic.AddInt64(&x, 1) // 原子操作
	// x=x+1
	// lock.Unlock()
	wg.Done()
}

func main() {
	wg.Add(100000)
	for i := 0; i < 100000; i++ {
		go add()
	}
	wg.Wait()
	fmt.Println(x)

	// 比较并交换
	// x = 200
	// fmt.Println(atomic.CompareAndSwapInt64(&x, 200, 100), x)
}
