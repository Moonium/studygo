package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.MaxFloat32)

	f1 := 1.23456
	fmt.Printf("%T\n", f1)
	f2 := float32(1.23456)
	fmt.Printf("%T\n", f2)
}
