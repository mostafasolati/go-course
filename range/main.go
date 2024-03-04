package main

import "fmt"

func main() {

	numbers := []int{1, 5, 6, 22}
	for _, n := range numbers {
		fmt.Println("value:", n)
	}

	scores := map[string]int{
		"sima":  17,
		"javid": 20,
		"elahe": 18,
		"mona":  20,
	}

	for k, v := range scores {
		fmt.Println("key:", k, "value:", v)
	}

	for i := range 10 {
		fmt.Println("number:", i)
	}

	alphabets := map[string]string{
		"a": "A",
		"b": "B",
		"c": "C",
	}

	for range 10 {
		for _, v := range alphabets {
			print(v, " ")
		}
		println()
	}
}
