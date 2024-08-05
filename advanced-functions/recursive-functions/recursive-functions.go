package main

import "fmt"

func fibonacci(n uint) uint {
	if n <= 1 {
		return n
	}

	return fibonacci(n - 1) + fibonacci(n - 2)
}

func main() {
	fmt.Println("recursive functions")


	for i := uint(1); i <= uint(12); i++ {
		fmt.Printf("fibonacci(%d) = %d\n", i, fibonacci(i))
	}
}