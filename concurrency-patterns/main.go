package main

import (
	"fmt"
	"sync"
	"time"
)

var c = make(chan int)
var registerdUsers = make(chan int, 3)
var num = 10

var wg sync.WaitGroup
var waitChan = make(chan int)

func waitForTask() {
	fmt.Println("wiating for a task")
	n := <-waitChan
	fmt.Println("Task", n, "done")
	wg.Done()
}

func registerUser(i int) {
	registerdUsers <- i
}

func double(n int) {
	c <- n * 2
}

var pool = make(chan string)

func worker(i int) {
	for {
		name := <-pool
		fmt.Println("goroutine:", i, "User", name, "registered")
		wg.Done()
	}
}

var max10 = make(chan int, 10)

func main() {

	go double(5)
	fmt.Println("double of 5 is:", <-c)

	for i := range num {
		go registerUser(i + 1)
	}

	for num > 0 {
		id := <-registerdUsers
		fmt.Println("User", id, "registerd")
		num--
	}

	wg.Add(1)
	go waitForTask()
	time.Sleep(time.Second * 2)
	waitChan <- 400
	wg.Wait()

	names := []string{"shiva", "maryam", "sina", "amirreza", "ali", "nima", "saeedeh", "sepdieh"}
	wg.Add(len(names))
	for i := range 5 {
		go worker(i)
	}

	for _, name := range names {
		pool <- name
	}
	wg.Wait()

	for i := range 20 {
		select {
		case max10 <- i:
		default:
			fmt.Println("request", i, "dropped")
		}
	}

	go func() {
		time.Sleep(2 * time.Second)
		c <- 5
	}()

	select {
	case n := <-c:
		fmt.Println("recieved number:", n)
	case <-time.After(time.Second):
		fmt.Println("Timeout")
	}
}
