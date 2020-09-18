package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// tcp client端

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("dial 127.0.0.1:20000 failed, err: %v", err)
		return
	}

	// var msg string
	reader := bufio.NewReader(os.Stdin)
	// if len(os.Args) < 2 {
	// 	msg = "Hello Internet!"
	// } else {
	// 	msg = os.Args[1]
	// }
	for {
		fmt.Print("请说话：")
		// fmt.Scanln(&msg)
		msg, _ := reader.ReadString('\n')
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}
	conn.Close()
}
