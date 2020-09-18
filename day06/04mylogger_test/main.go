package main

import (
	"github.com/Moonium/studygo/day06/mylogger"
)

var log mylogger.Logger // 声明一个全局的接口变量

// 测试自定义日志库

func main() {
	log = mylogger.NewConsoleLogger("Info")                             // 终端日志实例
	log = mylogger.NewFileLogger("Info", "./", "zhangsan.log", 10*1024) // 文件日志实例

	// log = mylogger.NewLogger("f")
	for {
		log.Debug("这是一条Debug日志")
		log.Info("这是一条Info日志")
		log.Warning("这是一条Warning日志")
		id := 10010
		name := "张三"
		log.Error("这是一条Error日志. id: %d name: %s", id, name)
		log.Fatal("这是一条Fatal日志")
		// time.Sleep(time.Second * 2)
	}
}
