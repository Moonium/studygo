package main

import (
	"fmt"
	"net"
	"strings"
)

// UDP server端
func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:]) // 接收数据
		if err != nil {
			fmt.Println("read udp failed, err:", err)
			return
		}
		// fmt.Println(data[:n])
		fmt.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)
		reply := strings.ToUpper(string(data[:n]))
		_, err = listen.WriteToUDP([]byte(reply), addr) // 发送数据
		if err != nil {
			fmt.Println("write to udp failed, err:", err)
			return
		}
	}
}
