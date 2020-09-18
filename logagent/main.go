package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/Moonium/studygo/logagent/conf"
	"github.com/Moonium/studygo/logagent/etcd"
	"github.com/Moonium/studygo/logagent/utils"

	"github.com/Moonium/studygo/logagent/kafka"
	"github.com/Moonium/studygo/logagent/taillog"
	"gopkg.in/ini.v1"
)

var (
	cfg = new(conf.AppConf)
)

// func run() {
// 	for {
// 		select {
// 		// 读取日志
// 		case line := <-taillog.ReadChan():
// 			// 发送到kafka
// 			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
// 		default:
// 			time.Sleep(time.Second)
// 		}
// 	}

// }

func main() {
	// 加载配置文件
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Println("load ini failed, err:", err)
		return
	}

	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.ChanMaxSize)
	if err != nil {
		fmt.Println("init kafka failed, err:", err)
		return
	}
	fmt.Println("init kafka success!")

	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Println("init etcd failed, err:", err)
		return
	}
	fmt.Println("init etcd success!")

	ipStr, err := utils.GetOutboundIP()
	if err != nil {
		panic(err)
	}
	etcdConfKey := fmt.Sprintf(cfg.EtcdConf.Key, ipStr)
	// 加载配置
	logEntryConf, err := etcd.GetConf(etcdConfKey)
	if err != nil {
		fmt.Println("etcd.GetConf failed, err:", err)
		return
	}
	fmt.Printf("get conf from etcd success! value: %v\n", logEntryConf)

	for index, value := range logEntryConf {
		fmt.Printf("index: %v, value: %v\n", index, value)
	}

	// 收集日志发往kafka
	taillog.Init(logEntryConf)
	// 派一个哨兵监视配置文件的变化
	newConfChan := taillog.NewConfChan()
	var wg sync.WaitGroup
	wg.Add(1)
	go etcd.WatchConf(etcdConfKey, newConfChan)
	wg.Wait()

	// err = taillog.Init(cfg.TaillogConf.FileName)
	// if err != nil {
	// 	fmt.Println("init taillog failed, err:", err)
	// 	return
	// }
	// fmt.Println("init taillog success!")

	// run()
}
