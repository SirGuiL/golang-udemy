package main

import "fmt"

func main() {
	countSum := sum(1, 2)

	fmt.Println(countSum)

	var function = func(txt string) {
		fmt.Println(txt)
	}

	function("Hello function")

	sum, sub := MathCalcs(1, 2)
	fmt.Println(sum, sub)

	sum2, _ := MathCalcs(5, 6)
	fmt.Println(sum2)
}

func MathCalcs(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	sub := n1 - n2

	return sum, sub
}

func sum(n1 int8, n2 int8) int8 {
	return n1 + n2
}