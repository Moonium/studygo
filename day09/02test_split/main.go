package main

import (
	"fmt"

	"github.com/Moonium/studygo/day09/splitstring"
)

func main() {
	ret := splitstring.Split("babcbef", "b")
	fmt.Printf("%#v\n", ret)
	ret2 := splitstring.Split("bbb", "b")
	fmt.Printf("%#v\n", ret2)
	ret3 := splitstring.Split("jklove", "b")
	fmt.Printf("%#v\n", ret3)
}
