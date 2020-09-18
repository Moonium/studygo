package taillog

import (
	"context"
	"fmt"

	"github.com/Moonium/studygo/logagent/kafka"
	"github.com/hpcloud/tail"
)

// 从日志文件收集日志的模块

// var tailObj *tail.Tail

// TailTask ...
type TailTask struct {
	path     string
	topic    string
	instance *tail.Tail
	// 为了实现退出
	ctx        context.Context
	cancelFunc context.CancelFunc
}

// NewTailTask ...
func NewTailTask(path, topic string) (tailObj *TailTask) {
	ctx, cancel := context.WithCancel(context.Background())
	tailObj = &TailTask{
		path:       path,
		topic:      topic,
		ctx:        ctx,
		cancelFunc: cancel,
	}
	tailObj.init()
	return
}

// Init ...
func (t *TailTask) init() {
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
	var err error
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("error->", err)
	}

	go t.run()
}

// // Init 初始化taillog
// func Init(fileName string) (err error) {
// 	// 设置config
// 	config := tail.Config{
// 		Location: &tail.SeekInfo{
// 			Offset: 0,
// 			Whence: 2,
// 		},
// 		Poll:      true,
// 		ReOpen:    true,
// 		MustExist: false,
// 		Follow:    true,
// 	}
// 	// 创建tail句柄
// 	tailObj, err = tail.TailFile(fileName, config)
// 	if err != nil {
// 		fmt.Println("error->", err)
// 		return
// 	}
// 	return
// }

// ReadChan 读取日志
func (t *TailTask) ReadChan() <-chan *tail.Line {
	return t.instance.Lines
}

func (t *TailTask) run() {
	for {
		select {
		case <-t.ctx.Done():
			fmt.Printf("tail task %s_%s 退出了...\n", t.path, t.topic)
			return
		case line := <-t.instance.Lines:
			// kafka.SendToKafka(t.topic, line.Text)
			kafka.SendToChan(t.topic, line.Text)
		}
	}
}
