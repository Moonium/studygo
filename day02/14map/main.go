package main

import "fmt"

func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil)        // 还没有初始化，没有在内存中开辟空间
	m1 = make(map[string]int, 10) // 要估算好map容量，避免在程序运行期间动态扩容
	m1["张三"] = 18
	m1["李四"] = 22

	fmt.Println(m1)
	fmt.Println(m1["张三"])
	fmt.Println(m1["王五"])
	v, ok := m1["王五"]
	if !ok {
		fmt.Println("查无此人")
	} else {
		fmt.Println(v)
	}

	// map的遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	for k := range m1 {
		fmt.Println(k)
	}
	for _, v := range m1 {
		fmt.Println(v)
	}

	delete(m1, "李四")
	fmt.Println(m1)
	delete(m1, "王五") // 删除不存在的key
}
