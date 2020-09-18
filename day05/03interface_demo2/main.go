package main

import "fmt"

// 接口的实现

type animal interface {
	move()
	eat(something string)
}

type cat struct {
	name string
	feet int8
}

func (c cat) move() {
	fmt.Println("走猫步~")
}

func (c cat) eat(food string) {
	fmt.Printf("%s吃%s！\n", c.name, food)
}

type chicken struct {
	name string
	feet int8
}

func (c chicken) move() {
	fmt.Println("鸡动！")
}

func (c chicken) eat(food string) {
	fmt.Printf("%s吃%s！", c.name, food)
}

func main() {
	var a1 animal
	bc := cat{
		name: "淘气",
		feet: 4,
	}
	a1 = bc
	fmt.Printf("%T", a1)
	a1.eat("鱼")

	kfc := chicken{
		feet: 2,
	}
	a1 = kfc
	fmt.Printf("%T", a1)
}
