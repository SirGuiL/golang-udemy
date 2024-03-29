package main

import "fmt"

func main() {
	var x1 int = 10
	var x2 int = x1

	x1++

	fmt.Println(x1, x2)

	// pointer is a memory reference
	var x3 int
	var pointer *int

	x3 = 10
	pointer = &x3

	fmt.Println(x3, pointer)
	fmt.Println(*pointer)

	x3 = 20
	fmt.Println(x3, pointer)
	fmt.Println(*pointer)
}