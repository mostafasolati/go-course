package main

import (
	"fmt"
	"sync"
	"time"
)

func setName(wg *sync.WaitGroup, ch chan<- string) {
	ch <- "Setareh"
	wg.Done()
}

func printName(wg *sync.WaitGroup, ch chan string) {
	fmt.Println("Hello", <-ch)
	wg.Done()
}

func heavyTask(i int, wg *sync.WaitGroup, q chan bool) {
	time.Sleep(time.Second)
	fmt.Println("Task", i, "Completed")
	wg.Done()
	<-q
}

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	ch := make(chan string)

	go setName(&wg, ch)
	go printName(&wg, ch)

	wg.Wait()

	ch2 := make(chan int, 2)
	ch2 <- 3
	ch2 <- 5
	// ch2 <- 7
	fmt.Println(<-ch2)

	const num = 5

	wg2 := sync.WaitGroup{}
	wg2.Add(num)
	queue := make(chan bool, 3)
	for i := range num {
		queue <- true
		go heavyTask(i, &wg2, queue)
	}
	wg2.Wait()

	messages := make(chan string, 3)
	go showMessages(messages)

	data := []string{
		"Hello",
		"my",
		"name",
		"is",
		"Mostafa",
	}
	for _, msg := range data {
		messages <- msg
		time.Sleep(time.Millisecond * 200)
	}
	// close(messages)

	time.Sleep(time.Second)
	fmt.Println("Finished!")

}

func showMessages(messages chan string) {
	// for msg := range messages {
	// 	fmt.Println(msg)
	// }
	for i := 0; i < 5; i++ {
		fmt.Println(<-messages)
	}
	fmt.Println("all data received")
}
