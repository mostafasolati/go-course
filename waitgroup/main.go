package main

import (
	"fmt"
	"sync"
)

func printLetters(wg *sync.WaitGroup) {
	for i := 'A'; i <= 'Z'; i++ {
		fmt.Print(string(i), " ")
	}
	wg.Done()
}

func printNumbers(wg *sync.WaitGroup) {
	for i := range 30 {
		fmt.Print(i, " ")
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go printLetters(&wg)
	go printNumbers(&wg)
	wg.Wait()
	fmt.Println("Finished")

	var a, b, c int
	var wg2 sync.WaitGroup
	wg2.Add(3)

	go func() {
		a = 1
		wg2.Done()
	}()

	go func() {
		b = 2
		wg2.Done()
	}()

	go func() {
		c = 3
		wg2.Done()
	}()
	wg2.Wait()
	fmt.Println(a + b + c)

}
