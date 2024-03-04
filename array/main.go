package main

import "fmt"

func main() {

	scores := [10]int{}
	scores[5] = 17
	fmt.Println(scores[5])

	primes := [5]int{1, 3, 5, 7, 11}
	fmt.Println(primes[4])

	for i := 0; i < 5; i++ {
		fmt.Print(primes[i], " ")
	}
	fmt.Println()

	fmt.Println("LEN:", len(primes))

	numbers := [20]int{}
	for i := 0; i < len(numbers); i++ {
		numbers[i] = i
	}
	fmt.Println(numbers[12])
	fmt.Println(numbers[19])
}
