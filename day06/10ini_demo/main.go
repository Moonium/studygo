package main

import "fmt"

// 需求分析

// MysqlConfig 结构体
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

func loadini(v interface{}) {

}

func main() {
	var mc MysqlConfig
	loadini(&mc)
	fmt.Println(mc.Address, mc.Port, mc.Username, mc.Password)
}
