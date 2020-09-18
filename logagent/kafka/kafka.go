package kafka

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

// 往kafka写日志的模块

var (
	client      sarama.SyncProducer
	logDataChan chan *logData
)

type logData struct {
	topic string
	data  string
}

// Init 初始化client
func Init(addrs []string, maxSize int) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 连接kafka
	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}

	logDataChan = make(chan *logData, maxSize)
	go sendToKafka()
	return
}

// SendToKafka 向kafka发送消息
func sendToKafka() {
	for {
		select {
		case ld:= <- logDataChan:
			msg := &sarama.ProducerMessage{}
			msg.Topic = ld.topic
			msg.Value = sarama.StringEncoder(ld.data)
			// 发送消息
			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				fmt.Println("send msg failed, err:", err)
				return
			}
			fmt.Printf("pid:%v offset:%v\n", pid, offset)
		default:
			time.Sleep(time.Millisecond * 50)
		}
		// 构造一个消息

	}
}

// SendToChan ...
func SendToChan(topic, data string) {
	msg := &logData{
		topic,
		data,
	}
	logDataChan <- msg
}
