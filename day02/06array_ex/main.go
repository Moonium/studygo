package main

import (
	"fmt"
)

func main() {
	a1 := [...]int{1, 3, 5, 7, 8}
	var sum int
	for _, v := range a1 {
		sum += v
	}
	fmt.Println(sum)

	for i := range a1 {
		for j := i; j < len(a1); j++ {
			if a1[i]+a1[j] == 8 {
				fmt.Println(i, j)
			}
		}
	}
}
