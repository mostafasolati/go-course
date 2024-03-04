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

func main() {
	u := User{
		FirstName: "Mostafa",
		LastName:  "Solati",
		Email:     "Mostafa.Solati@gmail.com",
		Age:       37,
	}

	// u := User{"Mostafa", "Solati", "mostafa.solati@gmail.com", 37}

	u.FirstName = "Mohammad Mostafa"
	fmt.Println(u.Email)

	var r rectangle
	r.width = 10
	r.height = 15
	fmt.Println(r)

	car := struct {
		model string
		year  int
		brand string
		price int
	}{
		model: "111",
		brand: "pride",
		price: 500000000,
	}

	fmt.Println(car.model, car.brand)
}
