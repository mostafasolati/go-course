package main

import (
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic Recover Reason:", r)
		}
	}()
	panic("we are going to die")
	fmt.Println("finish")
}
