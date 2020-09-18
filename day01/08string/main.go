package main

import (
	"fmt"
	"strings"
)

func main() {
	path := "\"D:\\code\\Go\\src\\github.com\\Moonium\\studygo\\day01\""
	fmt.Println(path)

	s := "I'm OK"
	fmt.Println(s)

	s2 := `
白日依山尽
	黄河入海流
		欲穷千里目
			更上一层楼
	`
	fmt.Println(s2)

	s3 := `D:\code\Go\src\github.com\Moonium\studygo\day01`
	fmt.Println(s3)

	fmt.Println(len(s3))

	name1 := "张三"
	name2 := "李四"
	fmt.Println(name1 + name2)

	ss := name1 + name2
	fmt.Println(fmt.Sprintf("%s%s%s", ss, name1, name2))

	ret := strings.Split(s3, "\\")
	fmt.Println(ret)

	fmt.Println(strings.Contains(ss, "王五"))
	fmt.Println(strings.Contains(ss, "张三"))

	fmt.Println(strings.HasPrefix(ss, "张"))
	fmt.Println(strings.HasSuffix(ss, "张"))

	s4 := "abcdeb"
	fmt.Println(strings.Index(s4, "c"))
	fmt.Println(strings.LastIndex(s4, "b"))

	fmt.Println(strings.Join(ret, "+"))
}
