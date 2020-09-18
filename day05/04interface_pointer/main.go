package main

// 使用值接收者和指针接受者的区别

import "fmt"

type animal interface {
	move()
	eat(something string)
}

type cat struct {
	name string
	feet int8
}

// // 方法使用值接收者
// func (c cat) move() {
// 	fmt.Println("走猫步~")
// }

// func (c cat) eat(food string) {
// 	fmt.Printf("%s吃%s！\n", c.name, food)
// }

// 方法使用指针接收者
func (c *cat) move() {
	fmt.Println("走猫步~")
}

func (c *cat) eat(food string) {
	fmt.Printf("%s吃%s！\n", c.name, food)
}

func main() {
	var a1 animal

	c1 := cat{"Tom", 4}
	c2 := &cat{"假老练", 4}

	a1 = &c1
	fmt.Println(a1)
	a1 = c2
	fmt.Println(a1)
}
