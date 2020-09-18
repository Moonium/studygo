package main

import (
	"bufio"
	"fmt"
	"os"
)

// 获取用户输入时如果有空格

func userScan() {
	var s string
	fmt.Print("请输入内容：")
	fmt.Scanln(&s)
	fmt.Println("您输入的内容是：", s)
}

func userBuio() {
	var s string
	fmt.Print("请输入内容：")
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	fmt.Println("您输入的内容是：", s)
}

func main() {
	// userScan()
	userBuio()

	// logger.Trace
	// logger.Debug()
	// logger.Warning()
	// logger.Info()
	// logger.Error()

	// 写日志
	fmt.Fprintln(os.Stdout, "这是一条日志记录！")
	fileObj, _ := os.OpenFile("./xxx.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	fmt.Fprintln(fileObj, "这是一条日志记录！")
}
