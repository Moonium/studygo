package main

import (
	"fmt"
	"path"
	"runtime"
)

func f() {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	// fmt.Println(pc)
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
	fmt.Println(file)
	fmt.Println(path.Base(file))
	fmt.Println(line)
}

func f1() {
	f()
}

func main() {
	f1()
}
