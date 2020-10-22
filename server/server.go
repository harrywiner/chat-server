package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

type Message struct {
	sender  int
	message string
}

func handleError(msg string, err error) {
	// TODO: all
	// Deal with an error event.
	if err != nil {
		fmt.Println("fugma at: ", msg)
	}
}

func acceptConns(ln net.Listener, conns chan net.Conn) {
	// TODO: all
	// Continuously accept a network connection from the Listener
	// and add it to the channel for handling connections.

	for {
		conn, err := ln.Accept()
		handleError("acceptConns", err)
		conns <- conn
	}

}

func handleClient(client net.Conn, clientid int, msgs chan Message) {
	// TODO: all
	// So long as this connection is alive:
	// Read in new messages as delimited by '\n's
	// Tidy up each message and add it to the messages channel,
	// recording which client it came from.
	reader := bufio.NewReader(client)
	for {
		msg, err := reader.ReadString('\n')

		// TODO this is gross
		if msg != "" {
			handleError("handleClient", err)

			newMessage := Message{
				sender:  clientid,
				message: msg}
			msgs <- newMessage
		}
	}
}

func main() {
	// Read in the network port we should listen on, from the commandline argument.
	// TODO CL argument
	// Default to port 8030
	portPtr := flag.String("port", ":8030", "port to listen on")
	flag.Parse()

	ln, err := net.Listen("tcp", *portPtr)
	handleError("Listen", err)
	fmt.Println("Listening on", *portPtr)

	//Create a channel for connections
	conns := make(chan net.Conn)
	//Create a channel for messages
	msgs := make(chan Message)
	//Create a mapping of IDs to connections
	clients := make(map[int]net.Conn)
	currentKey := 0

	//Start accepting connections
	go acceptConns(ln, conns)
	for {
		select {
		case conn := <-conns:
			//TODO Deal with a new connection
			// - assign a client ID
			// - add the client to the clients channel
			// - start to asynchronously handle messages from this client

			go handleClient(conn, currentKey, msgs)
			clients[currentKey] = conn
			fmt.Printf("Hooray! Welcome to user: %d \n", currentKey)
			currentKey++
		case msg := <-msgs:
			//TODO Deal with a new message
			// Send the message to all clients that aren't the sender

			fmt.Printf(">> %d: %s", msg.sender, msg.message)

			// iterate through the `clients`, and send a message to each one
			// if the clientid != msg.sender

			// this step is to get message sending working with client

			// next step is test multiple clients, then get their messages to send to eachother

			for i, client := range clients {
				if i != msg.sender {

					sender := msg.sender + 48

					toSend := ">> " + string(sender) + ": " + msg.message

					fmt.Fprintf(client, toSend)
				}
			}
		}
	}
}
