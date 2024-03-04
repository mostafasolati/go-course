package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:5050")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Succefully conntected to server on port 5050")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">> ")
		scanner.Scan()
		msg := scanner.Text()
		if msg == "exit" {
			return
		}
		fmt.Fprint(conn, msg)

	}
}
