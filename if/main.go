package main

import "fmt"

func main() {

	year := 1384

	if year >= 1361 && year <= 1370 {
		fmt.Println("60th")
	} else if year >= 1371 && year <= 1380 {
		fmt.Println("70th")
	} else if year >= 1381 && year <= 1390 {
		fmt.Println("80th")
	} else {
		fmt.Println("godzilaaaaa :D")
	}
}
