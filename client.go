package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func handleError(err error) {
	fmt.Println(err)
}

func read(conn net.Conn) {
	//TODO In a continuous loop, read a message from the server and display it.
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			handleError(err)
			return
		}
		fmt.Println(msg)
	}
}

func write(conn net.Conn) {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Enter text ->")
		msg, err := stdin.ReadString('\n')
		if err != nil {
			handleError(err)
			return
		}
		if msg == "/quit\n" {
			break
		}
		fmt.Fprintln(conn, msg)
	}
	//TODO Continually get input from the user and send messages to the server.
}

func main() {
	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()
	//TODO Try to connect to the server
	conn, _ := net.Dial("tcp", *addrPtr)
	//TODO Start asynchronously reading and displaying messages
	go read(conn)
	//TODO Start getting and sending user messages.
	write(conn)
}
