package main

import (
	"fmt"
	"net"
)

func main() {
	listner, err := net.Listen("tcp", "localhost:5050")
	if err != nil {
		panic(err)
	}
	defer listner.Close()
	fmt.Println("Succfully listening to port 5050")
	var i = 1
	for {
		conn, err := listner.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn, i)
		i++
	}
}

func handleConnection(conn net.Conn, i int) {
	fmt.Println("A new connection stablished...")
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("error recieveng data:", err)
			return
		}
		fmt.Println("Client", i, "Received:", string(buf[:n]))
	}
}
