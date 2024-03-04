package main

import (
	"fmt"
	"time"
)

func printLetters() {
	for i := 'A'; i <= 'Z'; i++ {
		fmt.Print(string(i), " ")
		// time.Sleep(time.Millisecond * 50)
	}
}

func printNumbers() {
	for i := range 30 {
		fmt.Print(i, " ")
		// time.Sleep(time.Millisecond * 50)
	}
}

func main() {
	go printLetters()
	go printNumbers()
	time.Sleep(2 * time.Second)
	fmt.Println("Finished")
}
