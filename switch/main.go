package main

import "fmt"

func main() {

	day := 1

	switch day {
	case 1:
		fmt.Println("shanbe")
	case 2:
		fmt.Println("1shanbe")
	case 3:
		fmt.Println("2shanbe")
	case 4:
		fmt.Println("3shanbe")
	case 5:
		fmt.Println("4shanbe")
	case 6:
		fmt.Println("5shanbe")
	case 7:
		fmt.Println("jome")
	default:
		fmt.Println("invalid day")
	}

	year := 1385

	switch {
	case year >= 1361 && year <= 1370:
		fmt.Println("60s")
	case year >= 1371 && year <= 1380:
		fmt.Println("70s")
	case year >= 1381 && year <= 1390:
		fmt.Println("80s")
	default:
		fmt.Println("godzilla")
	}

	food := "burger"
	price := 650000

	switch {
	case price >= 500000:
		fmt.Println("30% discount")
		fallthrough
	case food == "pizaa":
		fmt.Println("20% discount")
	case food == "burger":
		fmt.Println("10% discount")
	}
}
