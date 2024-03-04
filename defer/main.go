package main

import (
	"fmt"
)

func mytest() {
	defer fmt.Println("my test finished")
	fmt.Println("my test runned")
}

func main() {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")
	mytest()
	defer func() {
		defer fmt.Println("defer from closure")
		fmt.Println("closure exec1")
		fmt.Println("closure exec2")
	}()
}
