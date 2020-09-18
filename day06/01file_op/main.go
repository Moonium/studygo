package main

import (
	"fmt"
	"io"
	"os"
)

func f1() {
	var fileObj *os.File
	var err error
	fileObj, err = os.Open("./xxx.go")
	defer fileObj.Close()
	if err != nil {
		fmt.Printf("open file failed, err: %v", err)
		return
	}
}

func f11() {
	var fileObj *os.File
	var err error
	fileObj, err = os.Open("./main.go")
	defer fileObj.Close()
	if err != nil {
		fmt.Printf("open file failed, err: %v", err)
		return
	}
}

func f2() {
	// 打开要操作的文件
	fileObj, err := os.OpenFile("./sb.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("open file failed, err: %v", err)
		return
	}

	// 因为无法在文件中间插入内容，所以需要借助一个临时文件
	tempFile, err := os.OpenFile("./sb.tmp", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("create tmp file failed, error: %v", err)
		return
	}
	// defer tempFile.Close()

	// 读文件写入临时文件
	var ret [1]byte
	n, err := fileObj.Read(ret[:])
	if err != nil {
		fmt.Printf("read from file failed, err: %v", err)
		return
	}
	// 写入临时文件
	tempFile.Write(ret[:n])
	// 再写入要插入的内容
	tempFile.Write([]byte{'c'})
	// 紧接着把原文件后续的内容写入临时文件
	var x [1024]byte
	for {
		n, err := fileObj.Read(x[:])
		if err == io.EOF {
			tempFile.Write(x[:n])
			break
		}
		if err != nil {
			fmt.Printf("read from file failed, error: %v", err)
			return
		}
		tempFile.Write(x[:n])
	}
	// 原文件后续的也写入了临时文件中
	fileObj.Close()
	tempFile.Close()
	os.Rename("./sb.tmp", "./sb.txt")

	// fmt.Println(string(ret[:n]))

	// fileObj.Seek(1, 0) // 光标移动到b
	// fileObj.Write([]byte{'c'})

}

func main() {
	f2()
}
