package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		time.Sleep(time.Second)
		fmt.Println(i)
		i++
	}

	fmt.Println(i)

	for j := 0; j < 10; j++ {
		time.Sleep(time.Second)
		fmt.Println(j)
	}

	names := [5]string{"Guilherme", "Hugo", "Fagner", "Yuri Alberto", "Rodrigo Garro"}

	for index, name := range names {
		fmt.Println(index, name)
	}

	// only names
	for _, name := range names {
		fmt.Println(name)
	}

	// only index
	for index := range names {
		fmt.Println(index)
	}

	// loop in word
	for index, str := range "WORD" {
		fmt.Println(index, str)
	}

	// loop in map
	user := map[string]string {
		"name": "Guilherme",
		"city": "Guarulhos, SP",
	}

	for key, value := range user {
		fmt.Println(key, value)
	}
}