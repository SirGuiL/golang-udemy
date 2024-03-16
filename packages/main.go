package main

import (
	"fmt"
	"module/assistant"

	// external
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hello World!")

	assistant.Write()

	error := checkmail.ValidateFormat("guibiel-10@hotmail.com")
	fmt.Println(error)
}