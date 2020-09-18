package main

import (
	"fmt"

	"github.com/hpcloud/tail"
)

func main() {
	var (
		line *tail.Line
		ok   bool
	)
	// log文件名
	fileName := "./my.log"
	// 设置config
	config := tail.Config{
		Location: &tail.SeekInfo{
			Offset: 0,
			Whence: 2,
		},
		Poll:      true,
		ReOpen:    true,
		MustExist: false,
		Follow:    true,
	}
	// 创建tail句柄
	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("error->", err)
		return
	}
	for {
		// 通过管道获取到每条行数据
		line, ok = <-tails.Lines
		// fmt.Println("走这里了")
		if !ok {
			fmt.Println("tail file close, fileName:", tails.Filename)
			continue
		}
		fmt.Println("line:", line.Text)
	}

}
