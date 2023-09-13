package main

import (
	"errors"
	"fmt"
)

func main() {

	var num1 int = 10
	var ptr1 *int = &num1

	fmt.Println(ptr1, *ptr1)

	incr(num1) // call by value

	fmt.Println(num1)

	incrp(&num1) // pass or call by reference
	fmt.Println(num1)

	incrp(ptr1)
	fmt.Println(num1)

	ptr2 := new(int)   // var ptr2 *int , but
	fmt.Println(*ptr2) // type inference work
	*ptr2 = 200

	slice1 := make([]int, 3) // length 3 and cap 3

	var slice2 []int // declaration of the slice
	// two options.. either you can use make to assign memory
	// or you can use append to assign memory to the slice
	slice2 = append(slice2, 10)

	slice3 := []int{} // no memory
	//slice5 := make([]int, 0)
	//slice5[0] = 100

	var slice6 []int

	if slice3 == nil {
		fmt.Println("slice 3 is nil")
	}

	var slice4 []int // just declared

	if slice4 == nil {
		fmt.Println("slice 4 is nil")
	}

	fmt.Println(slice1, slice2, slice3, slice4)

	fmt.Println("Slice-1", len(slice1), cap(slice1))
	fmt.Println("Slice-2", len(slice2), cap(slice2))
	fmt.Println("Slice-3", len(slice3), cap(slice3))
	fmt.Println("Slice-4", len(slice4), cap(slice4))

	fmt.Println("Sum of slice1", sum(slice1))
	fmt.Println("Sum of slice6", sum(slice6))

	slice1[0] = 100
	slice1[1] = 98
	slice1[2] = 108
	fmt.Println("max of slice1", max(slice1)) // 108
	fmt.Println("max of slice6", max(slice6)) // [-23,-2,-4,-9,0]

	arr1 := [3]int{10, 11, 12}

	var arr2 [3]int // 0,0,0

	arr3 := [...]int{10, 11, 12, 13}

	fmt.Println(arr1, arr2, arr3)

	// sumArray(arr1)
	// sumArray(arr3)

	// arr5 := arr3[2:4]

	sum(arr1[:])
	sum(arr3[:])
	sum(arr3[2:4])
}

// difference between new and make
//

func incr(num int) {
	num++
}

func incrp(num *int) {
	*num++
}

/*

SliceHeader{
	HeaderPtr --> Pointer
	Len
	Cap -> How many elaments it can accomidate
}

// var slice1 []int
// slice2:=make([]int,0)
// slice3:=make([]int,0,10)
*/

// type Person struct {
// 	A         *string
// 	Apartment int
// 	People    int
// }

// var p1 *Person
// p2:= &Person{Apartment:101,People:10}
func sumArray(arr [3]int) int {
	sum := 0
	// for _, v := range slice {
	// 	sum += v
	// }

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func sum(slice []int) int {
	sum := 0
	// for _, v := range slice {
	// 	sum += v
	// }

	for i := 0; i < len(slice); i++ {
		sum += slice[i]
	}

	return sum
}

func max(slice []int) int {
	var big int

	//if len(slice) > 0 {
	if slice != nil && len(slice) > 0 {
		big = slice[0]
	} // no need to check nil when using len
	//}

	for i := 1; i < len(slice); i++ {
		if slice[i] > big {
			big = slice[i]
		}
	}

	return big
}

func maxE(slice []int) (int, error) {
	var big int

	if slice == nil {
		return 0, errors.New("nil slice")
	}
	//if len(slice) > 0 {
	if slice != nil && len(slice) > 0 {
		big = slice[0]
	} // no need to check nil when using len
	//}

	for i := 1; i < len(slice); i++ {
		if slice[i] > big {
			big = slice[i]
		}
	}

	return big, nil
}

// 1.10
// 1.20
