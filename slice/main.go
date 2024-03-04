package main

import "fmt"

func main() {
	users := make([]string, 0, 16)
	fmt.Println("before:")
	fmt.Println("slice:", users, "cap:", cap(users), "len:", len(users))
	users = append(users, "ali", "sima")
	fmt.Println("after:")
	fmt.Println("slice:", users, "cap:", cap(users), "len:", len(users))
	users = append(users, "sina")
	fmt.Println("slice:", users, "cap:", cap(users), "len:", len(users))
	users = append(users, "javad", "omid")
	fmt.Println("slice:", users, "cap:", cap(users), "len:", len(users))

	fmt.Println(users[2:5])
	fmt.Println(users[:3])
	fmt.Println(users[3:])
	fmt.Println(users[:], users)

	users2 := users[:2]
	fmt.Println(users2)
	users2[0] = "rezvan"
	fmt.Println(users2)
	fmt.Println(users)

	users3 := make([]string, 10)
	n := copy(users3, users)
	fmt.Println(users3)
	users3[0] = "samaneh"
	fmt.Println(users3)
	fmt.Println(users)
	fmt.Println(n, "items copied")
	clear(users)
	fmt.Println(users, "len:", len(users))

	numbers := []int{1, 2, 3, 5, 11}
	fmt.Println(numbers, len(numbers))
	clear(numbers)
	fmt.Println(numbers, len(numbers))
	numbers = make([]int, 0)
	fmt.Println(numbers, len(numbers))

}
