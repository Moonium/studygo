package main

import "fmt"

// 类型断言

func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	fmt.Println(str)
	if !ok {
		fmt.Println("猜错了")
	} else {
		fmt.Printf("传进来的是字符串：%s", str)
	}
}

func assign2(a interface{}) {
	fmt.Printf("%T\n", a)
	switch t := a.(type) {
	case string:
		fmt.Println("传进来的是字符串：", t)
	case int:
		fmt.Println("传进来的是整数：", t)
	case bool:
		fmt.Println("传进来的是布尔型：", t)
	default:
		fmt.Println("不支持的数据类型！")
	}

}

func main() {
	assign(100)

	assign2(100)
	assign2("100")
	assign2(true)
	assign2(nil)
}
