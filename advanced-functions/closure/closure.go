package main

import "fmt"

func closure() func() {
	text := "Inside closure function"

	return func() {
		fmt.Println(text)
	}
}

func main() {
	text := "Inside main function"
	fmt.Println(text)

	newFunction := closure()
	newFunction()
}