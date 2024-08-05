package main

import "fmt"

func revertSignal(number int) int {
	return -number
}

func revertSignalWithPointer(number *int) {
	*number = -*number
}

func main() {
	number := -10
	fmt.Println(number)
	reverted := revertSignal(number)
	fmt.Println(reverted)

	revertSignalWithPointer(&number)
	fmt.Println(number)
}