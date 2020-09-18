package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func writeDemo1() {
	fileObj, err := os.OpenFile("./xxx.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open file failed, err: %v", err)
		return
	}

	fileObj.Write([]byte("write file succeeded!\n"))

	fileObj.WriteString("这个问题解释不了！\n")
	fileObj.Close()
}
func writeDemo2() {
	fileObj, err := os.OpenFile("./xxx.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open file failed, err: %v", err)
		return
	}

	defer fileObj.Close()
	// 创建一个写的对象
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("Hello张三！\n") // 写到缓存中
	wr.Flush()                   // 将缓存中的内容写入文件
}

func writeDemo3() {
	str := "Hello张三！\n"
	err := ioutil.WriteFile("./xxx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Printf("open file failed, err: %v", err)
		return
	}
}

func main() {
	// writeDemo2()
	writeDemo3()
}
