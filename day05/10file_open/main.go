package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//打开文件
func readFromFile() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, error: %v", err)
		return
	}

	defer fileObj.Close()

	// var tmp = make([]byte, 128)
	var tmp [128]byte

	for {
		n, err := fileObj.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("读完了！")
			return
		}

		if err != nil {
			fmt.Printf("read from file failed, error: %v", err)
			return
		}
		fmt.Printf("读了%d个字节", n)
		fmt.Println(string(tmp[:n]))
		if n < 128 {
			return
		}
	}
}

func readFromFilebyBufio() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, error: %v", err)
		return
	}

	defer fileObj.Close()

	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read line failed, error: %v", err)
			return
		}
		fmt.Println(line)
	}
}

func readFromFilebyIoutil() {
	ret, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("read line failed, error: %v", err)
		return
	}
	fmt.Println(string(ret))
}

func main() {
	// readFromFile()
	readFromFilebyBufio()
}
