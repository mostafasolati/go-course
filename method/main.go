package main

import "fmt"

type (
	User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	rectangle struct {
		width  int
		height int
	}
)

type number int

func (s number) double() number {
	return s * 2
}

func (r rectangle) Primeter() int {
	return (r.width + r.height) * 2
}

func (r rectangle) Area() int {
	return r.width * r.height
}

func FullName(s User) string {
	return s.FirstName + " " + s.LastName
}

func (s User) FullName() string {
	return s.FirstName + " " + s.LastName
}

func main() {
	u := User{
		FirstName: "Mostafa",
		LastName:  "Solati",
		Email:     "Mostafa.Solati@gmail.com",
		Age:       37,
	}
	fmt.Println(u.FullName())
	fmt.Println(FullName(u))

	var r rectangle
	r.width = 10
	r.height = 15
	fmt.Println(r)

	fmt.Println("Area:", r.Area())
	fmt.Println("Primeter:", r.Primeter())

	var n number
	n = 5
	n = n.double()
	fmt.Println(n)

}
