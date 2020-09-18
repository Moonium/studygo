package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// tcp server端

func processConn(conn net.Conn) {
	defer conn.Close()
	var tmp [128]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Printf("read from connn failed, err: %v", err)
			return
		}
		fmt.Println(string(tmp[:n]))

		fmt.Print("请回复：")
		// fmt.Scanln(&msg)
		msg, _ := reader.ReadString('\n')
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("start server failed, err: %v", err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept failed, err: %v", err)
			return
		}
		go processConn(conn)
	}
}
