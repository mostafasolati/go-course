package main

import "fmt"

type number = int
type mathFunc func(int) int
type testFunc func(a, b, c int) func(string)

func apply(n int, f mathFunc) {
	result := f(n)
	fmt.Println(result)
}

func double(n number) number {
	return n * 2
}

func main() {
	var n int = 5
	fmt.Println(double(n))

	var f float64 = 3.14
	fmt.Println(int(f))

}
