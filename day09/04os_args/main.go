package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	fmt.Printf("%T\n%#v\n", os.Args, os.Args)
	fmt.Println(os.Args[0], os.Args[2])
}
