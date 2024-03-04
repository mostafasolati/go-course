package main

import "fmt"

type number int

func (n *number) double() {
	*n *= 2
}

type User struct {
	FirstName string
	LastName  string
	FullName  string
}

func (u *User) CalculateFullName() {
	u.FullName = u.FirstName + " " + u.LastName
}

func change(s []int) {
	s[0] = 1000
}

func main() {

	var n int = 5
	var p *int
	p = &n
	*p = 7
	fmt.Println("Value:", n)
	fmt.Println("Address:", p)
	fmt.Println("Value of Address:", *p)

	var num number
	num = 5
	num.double()
	fmt.Println(num)

	u := User{FirstName: "Mostafa", LastName: "Solati"}
	u.CalculateFullName()
	fmt.Println("FullName:", u.FullName)

	s := []int{1, 2, 3}
	change(s)
	fmt.Println(s)

}
