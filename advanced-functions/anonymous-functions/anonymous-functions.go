package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello World")
	}()

	func(str string) {
		fmt.Println(str)
	}("Hello World outside of the function")

	funcReturn := func(str string) string {
		return fmt.Sprintf("(Using Sprintf) Hello World %s", str)
	}("outside of the function")

	fmt.Println(funcReturn)
}