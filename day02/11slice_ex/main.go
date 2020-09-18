package main

import (
	"fmt"
	"sort"
)

func main() {
	var a = make([]string, 5, 10)
	// fmt.Println(a)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)
	fmt.Println(cap(a))

	var a1 = [...]int{3, 7, 8, 9, 1}
	sort.Ints(a1[:])
	fmt.Println(a1)
}
