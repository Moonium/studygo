package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "名字", "请输入名字")
	flag.Parse()
	fmt.Println(*name)

	// var name1 string
	// flag.StringVar(&name1, "name", "名字", "请输入名字")
	// flag.Parse()
	// fmt.Println(name1)

	fmt.Println(flag.Args())
	fmt.Println(flag.NArg())
	fmt.Println(flag.NFlag())
}
