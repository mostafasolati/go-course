package main

import (
	"fmt"
	"sync"
)

func main() {

	stock := 200
	sold := 0
	buyers := 0
	var wg sync.WaitGroup
	wg.Add(400)
	var mu sync.RWMutex

	for i := 1; i <= 400; i++ {

		go func(i int) {
			defer wg.Done()

			fmt.Println("buyer", i, "asks for stock")
			fmt.Println("seller says: we have", stock, "ps5")

			mu.RLock()
			if stock == 0 {
				mu.RUnlock()
				fmt.Println("out of sotck")
				return
			}
			mu.RUnlock()

			mu.Lock()
			sold++
			buyers++
			stock--
			mu.Unlock()
			fmt.Println("a ps5 purchased")

			//
		}(i)
	}

	wg.Wait()
	fmt.Println("----------------")
	fmt.Println("Unit sold:", sold)
	fmt.Println("Buyers:", buyers)
	fmt.Println("Stock:", stock)

}
