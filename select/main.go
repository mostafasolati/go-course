package main

import (
	"fmt"
	"time"
)

func main() {
	coffeeshopChannel := make(chan string)
	foodChannel := make(chan string)

	go func() {
		for i := 0; i < 1000; i++ {
			if i%2 == 0 {
				foodChannel <- "Burger"
				time.Sleep(time.Second)
			} else {
				coffeeshopChannel <- "Coffee"
				time.Sleep(time.Second * 3)
			}

		}
	}()

	for {
		select {
		case food := <-foodChannel:
			fmt.Println("A new food order:", food)
		case drink := <-coffeeshopChannel:
			fmt.Println("A new coffeeshop order:", drink)
		case <-time.After(time.Second * 2):
			fmt.Println("timeout!")
		}
	}

}
