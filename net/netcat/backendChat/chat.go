package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
)

type Client chan<- string

var (
	incomingClients = make(chan Client)
	leavingCLients  = make(chan Client)
	messages        = make(chan string)
)

var (
	chatHost = flag.String("h", "localhost", "host to connect to")
	chatPort = flag.Int("p", 3090, "port to connect to")
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	message := make(chan string)

	go MessageWriter(conn, message)

	clientName := conn.RemoteAddr().String()

	message <- fmt.Sprintf("Welcome to the server, your name %s\n", clientName)

	messages <- fmt.Sprintf("New client is here, name: %s\n", clientName)

	incomingClients <- message

	inputMessage := bufio.NewScanner(conn)

	for inputMessage.Scan() {
		messages <- fmt.Sprintf("%s: %s\n", clientName, inputMessage.Text())
	}

	leavingCLients <- message

	messages <- fmt.Sprintf("%s said bye\n", clientName)
}

func MessageWriter(conn net.Conn, messages <-chan string) {
	for message := range messages {
		fmt.Fprintf(conn, "%s\n", message)
	}
}

func Broadcast() {
	clients := make(map[Client]bool)

	for {
		select {
		case message := <-messages:
			for client := range clients {
				client <- message
			}

		case newClient := <-incomingClients:
			clients[newClient] = true

		case leavingCLient := <-leavingCLients:
			delete(clients, leavingCLient)
			close(leavingCLient)
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *chatHost, *chatPort))

	if err != nil {
		log.Fatal(err)
	}

	go Broadcast()

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Print(err)
			continue
		}

		go HandleConnection(conn)
	}
}
