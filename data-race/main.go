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
	for i := 1; i <= 400; i++ {

		go func() {

			defer wg.Done()

			fmt.Println("buyer", i, "asks for stock")
			fmt.Println("seller says: we have", stock, "ps5")
			if stock > 0 {
				fmt.Println("a ps5 purchased")
				sold++
				buyers++
				stock--
			} else {
				fmt.Println("sorry out of stock")
				return
			}
			//
		}()
	}

	wg.Wait()
	fmt.Println("----------------")
	fmt.Println("Unit sold:", sold)
	fmt.Println("Buyers:", buyers)
	fmt.Println("Stock:", stock)

}
