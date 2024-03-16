package main

import (
	"errors"
	"fmt"
)

func main() {
	// int8, int16, int32 (rune), int64
	// uint8 (byte), uint16, uint32, uint64

	// float32, float64
	// complex64, complex128

	// string
	// bool

	// error

	var err error = errors.New("Internal error")
	fmt.Println(err)

	fmt.Println("Hello World!")
}