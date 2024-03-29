package main

import "fmt"

func main() {
	fmt.Println(1 + 2)

	var arr1 [5]int

	fmt.Println(arr1)

	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	arr1[3] = 4
	arr1[4] = 5

	fmt.Println(arr1)

	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)


	slice = append(slice, 6)
	fmt.Println(slice)

	slice2 := arr2[1:3]
	fmt.Println(slice2)

	// intern arrays
	slice3 := make([]int, 5, 10)

	fmt.Println(slice3)
	
	fmt.Println(len(slice3)) // array length
	fmt.Println(cap(slice3)) // array capacity
}