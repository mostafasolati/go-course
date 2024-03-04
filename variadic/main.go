package main

import "fmt"

func sum(nums ...int) {
	var total int
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	fmt.Println("Sum:", total)
}

func main() {
	nums := []int{1, 2, 10, -1, -7}
	sum(nums...)
}
