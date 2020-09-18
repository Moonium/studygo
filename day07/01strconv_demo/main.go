package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := int32(2316)
	ret1 := string(i)
	fmt.Println(ret1)
	ret2 := fmt.Sprintf("%d", i)
	fmt.Printf("%#v %T\n", ret2, ret2)

	fmt.Println(strconv.Itoa(int(i)))

	str := "10000"
	ret3, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Printf("parse int failed, err: %v", err)
	}
	fmt.Printf("%#v %T\n", ret3, ret3)

	retInt, _ := strconv.Atoi(str)
	fmt.Printf("%#v %T\n", retInt, retInt)

	boolStr := "true"
	boolValue, _ := strconv.ParseBool(boolStr)
	fmt.Printf("%#v %T\n", boolValue, boolValue)
}
