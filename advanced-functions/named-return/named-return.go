package main

import "fmt"

func MathCalcs(n1, n2 int) (sum,  sub int) {
	sum = n1 + n2
	sub = n1 - n2

	return
}

func main() {
	sum, sub := MathCalcs(1, 2)
	fmt.Println(sum, sub)
}