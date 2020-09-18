package main

import "fmt"

type car interface {
	run()
}

type ferrari struct {
	brand string
}

func (f ferrari) run() {
	fmt.Printf("%s速度70迈~\n", f.brand)
}

type porsche struct {
	brand string
}

func (p porsche) run() {
	fmt.Printf("%s速度80迈~\n", p.brand)
}

func drive(c car) {
	c.run()
}

func main() {
	var f1 = ferrari{
		brand: "法拉利",
	}
	var p1 = porsche{
		brand: "保时捷",
	}
	drive(f1)
	drive(p1)
}
