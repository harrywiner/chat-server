package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func handleError(msg string, err error) {
	// TODO: all
	// Deal with an error event.
	if err != nil {
		fmt.Println("fugma at: ", msg)
	}
}

func read(conn *net.Conn) {
	//TODO In a continuous loop, read a message from the server and display it.

	reader := bufio.NewReader(*conn)
	for {
		msg, err := reader.ReadString('\n')

		handleError("Client Read", err)

		// TODO this is gross
		if msg != "" {
			fmt.Println(msg)
		}
	}
}

func write(conn *net.Conn) {
	//TODO Continually get input from the user and send messages to the server.

	stdin := bufio.NewReader(os.Stdin)
	//fmt.Print(">> Me: ")
	for {
		msg, err := stdin.ReadString('\n')

		handleError("Client Write", err)

		// TODO
		if msg != "nil" {
			fmt.Fprintf(*conn, msg)
			//fmt.Print(">> Me: ")
		}
	}

}

func main() {
	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()
	//TODO Try to connect to the server
	//TODO Start asynchronously reading and displaying messages
	//TODO Start getting and sending user messages.
	conn, _ := net.Dial("tcp", *addrPtr)

	fmt.Fprintf(conn, "Hello Boozer!\n")

	go write(&conn)

	go read(&conn)

	time.Sleep(time.Second * 120)
}
