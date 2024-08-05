package main

import "fmt"

func recoveryExecution() {
	if r := recover(); r != nil {
		fmt.Println("Recovered in panic", r)
	}
}

func studentIsApproved(n1, n2 float64) bool {
	defer recoveryExecution()
	average := (n1 + n2) / 2

	if average > 6 {
		return true
	} else if average < 6 {
		return false
	}

	panic("Average is equal to 6")
}

func main() {
	fmt.Println(studentIsApproved(6, 7))
	fmt.Println("passed")
	fmt.Println(studentIsApproved(6, 6))
	fmt.Println("passed")
}