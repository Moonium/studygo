package main

import "fmt"

// Go语言下的return函数不是原子操作，而是分为两步进行：
// 1.返回值赋值
// 2.执行真正的RET返回
// 函数中如果存在defer，那么defer执行是在上述两步之间

func f1() int {
	x := 5
	defer func() {
		x++ // 修改的是x不是返回值
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 // 返回值 = x
}

func f3() int {
	x := 5
	defer func() {
		x++ // 修改的是x不是返回值
	}()
	return x // 返回值 = y = x = 5
}

func f4() (x int) {
	defer func(x int) {
		x++ // 将x当做参数传入，改变的是函数中的副本
	}(x)
	return 5 // 返回值 = x = 5
}

func f5() (x int) {
	defer func(x int) int {
		x++
		return x
	}(x)
	return 5
}

// 传一个x的指针
func f6() (x int) {
	defer func(x *int) {
		*x++

	}(&x)
	return 5
}

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
	fmt.Println(f5())
	fmt.Println(f6())
}
