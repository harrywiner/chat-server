package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func read(conn *net.Conn) {
	reader := bufio.NewReader(*conn)
	msg, _ := reader.ReadString('\n')
	fmt.Println(">> ", msg)
}

func main() {
	stdin := bufio.NewReader(os.Stdin)
	for {
		// read flags from cli
		msg, _ := stdin.ReadString('\n')
		conn, _ := net.Dial("tcp", "127.0.0.1:8000")
		// send message
		fmt.Fprintf(conn, msg)

		// read callback message
		read(&conn)
	}
}
