package main

import (
	"fmt"

	"github.com/Moonium/studygo/log_transfer/conf"
	"github.com/Moonium/studygo/log_transfer/es"
	"github.com/Moonium/studygo/log_transfer/kafka"
	"gopkg.in/ini.v1"
)

func main() {
	var cfg = new(conf.LogTransferCfg)
	err := ini.MapTo(cfg, "./conf/cfg.ini")
	if err != nil {
		fmt.Printf("init cfg failed, err: %v", err)
		return
	}
	fmt.Printf("cfg: %v", cfg)

	err = es.Init(cfg.ESCfg.Address, cfg.ESCfg.ChanSize, cfg.ESCfg.Nums)
	if err != nil {
		fmt.Printf("init es config failed, err: %v", err)
		return
	}
	fmt.Println("init es success!")

	err = kafka.Init([]string{cfg.Kafka.Address}, cfg.Kafka.Topic)
	if err != nil {
		fmt.Printf("init kafka consumer failed, err: %v", err)
		return
	}

	select {}
}
