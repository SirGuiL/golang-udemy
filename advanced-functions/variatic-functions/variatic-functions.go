package main

import "fmt"

func sum(nums ...int) {
	total := 0
	
	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)
	sum(1, 2, 3, 4)
	sum(1, 2, 3, 4, 5)
	sum(1, 2, 3, 4, 5, 6)
	sum(1, 2, 3, 4, 5, 6, 7)
	sum(1, 2, 3, 4, 5, 6, 7, 8)
	sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}