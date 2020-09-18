package main

import "fmt"

func main() {
	n := 18
	fmt.Println(&n)
	p := &n
	fmt.Printf("%T\n", p)
	m := *p
	fmt.Println(m)
	fmt.Printf("%T\n", m)

	// a1 := [...]int{1, 2, 3}
	// fmt.Println(&a1[0])
	// s1 := a1[1:]
	// fmt.Println(a1)
	// fmt.Println(&a1[0])
	// fmt.Println(s1)
	// fmt.Println(&s1[0])

	var a1 *int // nil pointer
	fmt.Println(a1)
	var a2 = new(int) // new分配内存，返回对应类型的指针
	fmt.Println(a2)
	*a2 = 100
	fmt.Println(*a2)

}
