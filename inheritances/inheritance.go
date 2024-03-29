package main

import "fmt"

type people struct {
	name   string
	age    uint8
	height float32
}

type student struct {
	people
	college string
}

func main() {
	p1 := people{"Guilherme", 21, 1.75}
	s1 := student{p1, "USP"}
	fmt.Println(s1)
}