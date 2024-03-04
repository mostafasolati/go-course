package main

import "fmt"

func main() {

	// students := make(map[string]int)
	students := map[string]int{
		"masoud": 20,
		"reza":   18,
	}
	fmt.Println(students)
	fmt.Println("Reza:", students["reza"])
	fmt.Println(len(students))

	delete(students, "masoud")
	fmt.Println(students)

	score, ok := students["masoud"]
	fmt.Println(score, ok)

	clear(students)
	fmt.Println(students, len(students))
}
