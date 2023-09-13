package main

import (
	"fmt"
	"math/rand"
)

func main() {

	slice1 := make([]int, 5) // length 5 and cap 5

	for i := range slice1 {
		slice1[i] = rand.Intn(100)
	}

	slice2 := slice1 // shallow copy
	fmt.Println("initially")
	fmt.Println("slice1=", slice1)
	fmt.Println("slice2", slice2)

	slice2[1] = 9000
	fmt.Println("After changing-1")
	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2)

	slice1 = append(slice1, 200, 300, 400, 500)

	slice2[1] = 9999
	fmt.Println("After changing-2")
	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2)

	fmt.Println("Multiply With 2")
	slice3 := []int{10, 11, 12} // assign values directly
	//slice3 := make([]int, 3, 20)
	fmt.Println("slice3:", slice3)
	slice3 = MulX2(slice3, 2, 13, 14, 15, 16, 17, 18, 19, 20)
	fmt.Println("slice3:", slice3)
}

// take the slice , multiply each element of the slice with 2 and check any new elements
// to be added to the slice and multiply them with 2 as well
func MulX2(slice []int, mul int, nums ...int) []int {
	slice = append(slice, nums...)
	for i, v := range slice {
		slice[i] = v * mul
	}
	return slice
}
