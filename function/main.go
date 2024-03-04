package main

import "fmt"

var number int

func hello(name string) {
	fmt.Println("Hello ", name)
}

func sum(a, b int) int {
	return a + b
}

func math(a, b int) (int, int, int, int) {
	res1 := sum(a, b)
	res2 := a - b
	res3 := a * b
	res4 := a / b
	return res1, res2, res3, res4
}

func math2(a, b int) (res1 int, res2 int, res3 int, res4 int) {
	number := 1000
	res1 = a + b
	res2 = a - b
	res3 = a * b
	res4 = a / b
	fmt.Println("shadow number:", number)
	return
}

func main() {

	hello("Gopher")
	fmt.Println(sum(5, 7))
	fmt.Println("Number:", number)
	// x, y, z, t := math(5, 7)
	x, y, z, t := math2(5, 7)
	fmt.Println("Number:", number)
	fmt.Println("sum:", x, "sub:", y, "dot:", z, "div:", t)
}
