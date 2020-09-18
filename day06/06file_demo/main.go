package main

import (
	"fmt"
	"os"
)

// 1. 文件对象的类型
// 2. 获取文件对象的详细信息

func main() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err: %v", err)
		return
	}
	fmt.Printf("%T\n", fileObj)

	fileInfo, err := fileObj.Stat()
	if err != nil {
		fmt.Printf("get file info failed, err: %v", err)
		return
	}
	fmt.Printf("文件大小是：%d 字节\n", fileInfo.Size())
	fmt.Printf("文件名是：%s\n", fileInfo.Name())
}
