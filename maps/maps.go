package main

import "fmt"

func main() {
	user := map[string]string {
		"name": "Guilherme",
		"city": "São Paulo",
	}
	
	fmt.Println(user)

	user2 := map[string]map[string]string {
		"name": {
			"first": "Guilherme",
			"last": "Lima",
		},
		"address": {
			"state": "SP",
			"city": "São Paulo",
		},
	}

	fmt.Println(user2)
	delete(user2, "address")
	fmt.Println(user2)

	user2["course"] = map[string]string {
		"name": "Golang",
		"professor": "Otávio Augusto",
	}

	fmt.Println(user2)
}