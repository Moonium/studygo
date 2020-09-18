package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name" db:"name" ini:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		"张三",
		22,
	}

	// 序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("%#v\n", string(b))

	// 反序列化
	str := `{"name":"李四", "age":18}`
	var p2 person
	json.Unmarshal([]byte(str), &p2) // 传指针是为了能在函数中修改p2的值
	fmt.Printf("%#v\n", p2)
}
