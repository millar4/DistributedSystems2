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

func handleError(err error) {
	if err != nil {
		fmt.Println("There is an error")
	}

}

func acceptConns(ln net.Listener, conns chan net.Conn) {
	// TODO: all
	// Continuously accept a network connection from the Listener
	// and add it to the channel for handling connections.

	for {
		conn, err := ln.Accept()
		handleError(err)
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
		message, err := reader.ReadString('\n')
		if err != nil {
			handleError(err)
			break
		}

		var msg Message
		msg.sender = clientid
		msg.message = message

		msgs <- msg
	}
}

func main() {
	// Read in the network port we should listen on, from the commandline argument.
	// Default to port 8030
	portPtr := flag.String("port", ":8030", "port to listen on")
	flag.Parse()

	//TODO Create a Listener for TCP connections on the port given above.
	ln, err := net.Listen("tcp", *portPtr)
	handleError(err)
	//Create a channel for connections
	conns := make(chan net.Conn)
	//Create a channel for messages
	msgs := make(chan Message)
	//Create a mapping of IDs to connections
	clients := make(map[int]net.Conn)

	clientid := 0
	//Start accepting connections
	go acceptConns(ln, conns)
	for {
		select {
		case conn := <-conns:
			//TODO Deal with a new connection
			// - assign a client ID
			// - add the client to the clients map
			client := conn
			clients[clientid] = conn
			// - start to asynchronously handle messages from this client
			go handleClient(client, clientid, msgs)
			clientid += 1
		case msg := <-msgs:
			for i, client := range clients {
				if msg.sender != i {
					fmt.Fprintf(client, "\n%s", msg.message)
				}
			}
		}
	}
}
