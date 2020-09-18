package main

import "fmt"

// Go语言中标识符的首字母大写表示对外部可见（公有的）

// Dog 是一个狗的结构体
type Dog struct {
	name string
}

func newDog(name string) Dog {
	return Dog{
		name: name,
	}
}

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

// 使用值接收者：传拷贝进去
func (p person) newYear() {
	p.age++
}

// 指针接收者：传内存地址进去
func (p *person) realNewYear() {
	p.age++
}

func (p *person) dream() {
	fmt.Println("做梦")
}

// 方法是作用于特定类型的函数
// 接收者是调用该方法的具体的类型，多用类型名首字母小写表示
func (d Dog) wang() {
	fmt.Printf("%s在汪汪汪！\n", d.name)
}

func main() {
	d1 := newDog("旺财")
	d1.wang()

	p1 := newPerson("张三", 22)
	fmt.Println(p1.age)
	p1.newYear()
	fmt.Println(p1.age)
	p1.realNewYear()
	fmt.Println(p1.age)

	p1.dream()
}
