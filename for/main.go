package main

import "fmt"

func main() {

	for i := 10; i >= 0; i-- {
		fmt.Print(i, ",")
	}

	var counter int
	fmt.Print("\nzoj: ")
	for i := 0; i <= 10; i++ {
		counter++
		if i%2 == 0 {
			fmt.Print(i, ",")
		}
	}
	fmt.Println("\nCounter:", counter)

	fmt.Println("zoj new algorithm: ")
	counter = 0
	for i := 0; i <= 10; i += 2 {
		counter++
		fmt.Print(i, ",")
	}
	fmt.Println("\nCounter:", counter)

	// i = i + x --> i += x
	// t = t -x --> t -= x
	// t = t * x --> t *= x
	// t = t / x --> t /= x

	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Print(i*j, "\t")
		}
		fmt.Println()
	}

	for i := 1; i <= 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	for i := 10; i >= 0; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	i := 0
	for {
		i++
		if i == 5 {
			continue
		}
		fmt.Print(i, ",")
		if i > 9 {
			break
		}
	}

	fmt.Println()
}
