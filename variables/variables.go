package main

import "fmt"

func main() {
	var a string = "initial"
	b := "short"

	fmt.Println(a, b)

	var (
		c int = 1
		d int = 2
	)

	fmt.Println(c, d)
}