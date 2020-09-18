package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fileObj, err := os.OpenFile("./xx.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failes, err: %v", err)
		return
	}
	log.SetOutput(fileObj)
	for {
		log.Println("这是一条测试的日志~")
		time.Sleep(time.Second * 3)
	}
}
