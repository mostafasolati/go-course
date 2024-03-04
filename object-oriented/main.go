package main

import "fmt"

type A interface{}
type B interface{}
type C interface {
	A
	B
}

type User struct {
	FirstName string
	LastName  string
	Age       int
}

func (u User) FullName() string {
	return u.FirstName + " " + u.LastName
}

type Customer struct {
	user   User
	Creidt int
}

type Courier struct {
	user   User
	Salary int
}

type Singer struct {
	genere string
	Age    int
	User
}

func main() {
	customer := Customer{
		Creidt: 500000,
		user: User{
			FirstName: "Ramin",
			LastName:  "Javadi",
			Age:       40,
		},
	}

	courier := Courier{
		Salary: 50000000,
		user: User{
			FirstName: "Javad",
			LastName:  "Ramini",
		},
	}

	singer := Singer{
		genere: "pop",
		User: User{
			FirstName: "Ebrahim",
			LastName:  "Hamedi",
		},
	}

	fmt.Println(customer.Creidt)
	fmt.Println(courier.user.FirstName)

	fmt.Println(customer.user.FullName())
	fmt.Println(courier.user.FullName())

	singer.FirstName = "Ebi"
	singer.Age = 80
	singer.User.Age = 90
	fmt.Println(singer.FullName())
	fmt.Println(singer)

}
