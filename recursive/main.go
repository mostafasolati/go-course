// 5! = 5 x 4 x 3 x 2 x 1 = 120

package main

import "fmt"

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

/*
factorial(5) : 5 * 4 * 3 * 2 * 1
factorial(4) : 4 * 3 * 2 * 1
factorial(3) : 3 * 2 * 1
factorial(2) : 2 * 1
factorial(1) : 1
*/

// 1 1 2 3 5 8 13

func fib(a, b, n int) {
	if n <= 0 {
		return
	}
	fmt.Println(a, "+", b, "=", a+b)
	fib(b, a+b, n-1)
}

func main() {
	fmt.Println(factorial(5))
	fib(1, 1, 5)
}
