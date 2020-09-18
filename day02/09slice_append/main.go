package main

import "fmt"

func main() {
	s1 := []string{"北京", "上海", "广州"}
	fmt.Printf("s1 =%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, "深圳")
	fmt.Printf("s1 =%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	// fmt.Println(s1)
	s1 = append(s1, "杭州")
	fmt.Printf("s1 =%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	ss := []string{"武汉", "成都", "西安", "苏州"}
	s1 = append(s1, ss...)
	fmt.Printf("s1 =%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
}
