package main

import "fmt"

func f(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}

func stair(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return stair(n-1) + stair(n-2)
}

func main() {
	fmt.Println(f(5))
	fmt.Println(stair(10))
}
