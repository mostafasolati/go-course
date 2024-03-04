package main

import "fmt"

func test() {

}

func NewCounter(n int) func() {

	var num, score int
	num = n
	f := func() {
		score++
		fmt.Println("Player", num, "Score:", score)
	}
	return f
}

func main() {

	var c int
	counter := func() int {
		c++
		return c
	}

	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	var score1, score2 int
	player1 := func() int {
		score1++
		return score1
	}
	player2 := func() int {
		score2++
		return score2
	}

	fmt.Println("player1:", player1())
	fmt.Println("player1:", player1())
	fmt.Println("player2:", player2())
	fmt.Println("player1:", player1())

	for i := 1; i <= 10; i++ {
		f := NewCounter(i)
		f()
	}

	player300 := NewCounter(300)
	player400 := NewCounter(400)
	player500 := NewCounter(500)

	player300()
	player400()
	player500()
	player500()

	func(name string) {
		fmt.Println("Hello", name)
	}("Sara")

	fmt.Println(apply(5, double))
	fmt.Println(apply(5, square))
	fmt.Println(apply(5, func(n int) int {
		return n + 1
	}))

	ShowMenu(Pizza)
	fmt.Println()
	ShowMenu(Burger)

}

func double(n int) int {
	return n * 2
}

func square(n int) int {
	return n * n
}

func apply(x int, f func(int) int) int {
	return f(x)
}

func Pizza() (string, int) {
	return "Pizza", 350000
}

func Burger() (string, int) {
	return "Burger", 200000
}

func ShowMenu(f func() (string, int)) {
	name, price := f()
	fmt.Println("Name:\t", name)
	fmt.Println("--------------")
	fmt.Println("Price:", price)
}
