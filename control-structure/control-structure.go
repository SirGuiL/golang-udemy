package main

import "fmt"

func main() {
	number := 10

	if number > 10 {
		fmt.Println("Number is greater than 10")
	} else if number < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is equal to 10")
	}

	if number2 := number; number2 > 10 {
		fmt.Println("Number is greater than 10")
		fmt.Println(number2)
	} else {
		fmt.Println(number2)
	}
}