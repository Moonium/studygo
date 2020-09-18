package main

import (
	"fmt"
)

const pi = 3.1415926

const (
	n1 = 100
	n2
	n3
)

const (
	a1 = iota
	a2
	a3
)

const (
	statusOK = 200
	notFound = 404
)

const b1 = iota
const c1 = iota
const b2 = iota

const (
	d1 = iota
	d2
	e1
	d3
)

const (
	f1 = iota
	f2 = 100
	f3 = iota
	f4
)

const (
	g1, g2 = iota + 1, iota + 2

	g3, g4 = iota + 1, iota + 2
)

const (
	_  = iota
	kb = 1 << (10 * iota)
	mb = 1 << (10 * iota)
	gb = 1 << (10 * iota)
	tb = 1 << (10 * iota)
	pb = 1 << (10 * iota)
)

func main() {
	fmt.Println(statusOK / 5)
	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)
	fmt.Println("n3:", n3)

	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println("a3:", a3)

	fmt.Println("b1:", b1)
	fmt.Println("b2:", b2)

	fmt.Println("d1:", d1)
	fmt.Println("d2:", d2)
	fmt.Println("d3:", d3)

	fmt.Println("f1:", f1)
	fmt.Println("f2:", f2)
	fmt.Println("f3:", f3)
	fmt.Println("f4:", f4)

	fmt.Println("g1:", g1)
	fmt.Println("g2:", g2)
	fmt.Println("g3:", g3)
	fmt.Println("g4:", g4)

	fmt.Println(kb)
	fmt.Println(mb)
	fmt.Println(gb)
	fmt.Println(tb)
	fmt.Println(pb)
}
