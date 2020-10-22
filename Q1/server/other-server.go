package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn *net.Conn) {
	// read message from conn
	reader := bufio.NewReader(*conn)
	msg, _ := reader.ReadString('\n')

	// print msg
	fmt.Println(">> ", msg)

	// Send back received message
	fmt.Fprintf(*conn, "Received\n")
}

func main() {
	ln, _ := net.Listen("tcp", "127.0.0.1:8000")
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("oh fug")
		}

		handleConnection(&conn)

	}
}
