package main

import "fmt"

type user struct {
	name string
	age uint8
	address address
}

type address struct {
	state string
	city string
}

func main() {
	fmt.Println("Struct file")

	var user1 user
	user1.name = "Guilherme"
	user1.age = 21

	fmt.Println(user1)

	user2 := user{"Guilherme", 21, address{"SP", "São Paulo"}}
	fmt.Println(user2)

	user3 := user{name: "Guilherme"}
	fmt.Println(user3)

	user4 := user{age: 21}
	fmt.Println(user4)
}