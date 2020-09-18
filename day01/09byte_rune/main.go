package main

import (
	"fmt"
)

func main() {
	s := "Hello张三"
	n := len(s)
	fmt.Println(n)

	for _, c := range s {
		fmt.Printf("%c\n", c)
	}

	s2 := "白萝卜"
	s3 := []rune(s2)
	s3[0] = '红'
	fmt.Println(string(s3))

	c1 := "红"
	c2 := '红'
	fmt.Printf("c1:%T c2:%T\n", c1, c2)

	c3 := "H"
	c4 := byte('H')
	fmt.Printf("c3:%T c4:%T\n", c3, c4)
	fmt.Printf("%d\n", c4)

	n1 := 10
	var f float64
	f = float64(n1)
	fmt.Println(f)
	fmt.Printf("%T\n", f)
}
