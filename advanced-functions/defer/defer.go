package main

import "fmt"

func deferFunc1() {
	fmt.Println("Hello World")
}

func deferFunc2() {
	fmt.Println("Hello World 2")
}

func main() {
	defer deferFunc1()
	deferFunc2()
}