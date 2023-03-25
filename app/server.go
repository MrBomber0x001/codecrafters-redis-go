package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	// Uncomment this block to pass the first stage
	// "net"
	// "os"
)

var (
	wg sync.WaitGroup
)

func main() {
	wg.Add(1)
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go webWorker(conn)
	}
}

func webWorker(conn net.Conn) {
	defer conn.Close()

	buff := make([]byte, 1024)

	for {
		if _, err := conn.Read(buff); err != nil {
			fmt.Println("erro reading from client", err.Error())
			os.Exit(1)
		}
		conn.Write([]byte("+PONG\r\n"))
	}
}
