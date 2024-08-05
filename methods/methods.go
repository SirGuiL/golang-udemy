package main

import "fmt"

type user struct {
	name string
	age uint8
}

func (u user) save() {
	fmt.Printf("User saved with name %s and age %d\n", u.name, u.age)
}

func (u user) ofLegalAge() bool {
	return u.age >= 18
}

func (u *user) birthday() {
	u.age++
}

func main() {
	user1 := user{"Guilherme", 21}

	fmt.Println(user1)

	user1.save()
	ofLegalAge := user1.ofLegalAge()
	user1.birthday()

	fmt.Println(ofLegalAge)
	fmt.Println(user1)
}