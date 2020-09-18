package main

import (
	"fmt"
	"strings"
)

func makeSuffixfunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	txtFunc := makeSuffixfunc(".txt")
	jpgFunc := makeSuffixfunc(".jpg")
	fmt.Println(txtFunc("test"))
	fmt.Println(jpgFunc("test"))
}
