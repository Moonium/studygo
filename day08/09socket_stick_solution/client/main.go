package main

import (
	"fmt"
	"net"

	proto "github.com/Moonium/studygo/day08/09socket_stick_solution/protocol"
)

// socket_stick/client/main.go

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		// 调用协议编码
		b, _ := proto.Encode(msg)
		conn.Write([]byte(b))
		// time.Sleep(time.Second)
	}
}
