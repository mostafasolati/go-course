package main

import "fmt"

func main() {
	var name string = "Mostafa"
	var movie string = "Godfather"
	var score float64 = 9.5
	var rottenTomato int = 87

	name = "Parmis"
	movie = "Barbie"
	score = 7.5

	fmt.Println(name, "loves", movie, "movie")
	fmt.Println("it's imdb score is:", score)
	fmt.Println("it's rotten tomato scores is:", rottenTomato)

	fmt.Println(`Hello
	
	
	
World`)

}
