package main

import (
	"fmt"
)

func main() {
	var (
		a = 5
		b = 2
	)
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	a++ // 单独语句，不能放在等号右边给变量赋值
	b--
	fmt.Println(a)

	fmt.Println(a == b) // Go语言是强类型，相同类型的变量才能相互比较
	fmt.Println(a != b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
	fmt.Println(a > b)
	fmt.Println(a < b)

	age := 20
	if age >= 18 && age <= 60 {
		fmt.Println("上班")
	} else {
		fmt.Println("不用上班")
	}

	if age < 18 || age > 60 {
		fmt.Println("不用上班")
	} else {
		fmt.Println("上班")
	}

	isMarried := false
	fmt.Println(isMarried)
	fmt.Println(!isMarried)

	fmt.Println(5 & 2)
	fmt.Println(5 | 2)
	fmt.Println(5 ^ 2)
	fmt.Println(5 << 2)
	fmt.Println(5 >> 2)

	// var m = int8(1)
	// fmt.Println(m << 10)

	var x int
	x = 10
	x++
	x -= 2
	x *= 3
	x /= 4
	x %= 2

	x <<= 2
	x &= 2
	x |= 2
	x ^= 4
	fmt.Println(x)
}
