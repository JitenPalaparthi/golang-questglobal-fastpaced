package main

import (
	"fmt"
	"math/rand"
)

func main() {

	si1 := make(sliceInt, 10) // can use make
	si1.Feed()
	si1 = append(si1, 12323123)
	sum := si1.Sum()
	fmt.Println(sum)

	slice2 := make([]int, 10)
	copy(slice2, si1)

	fmt.Println(slice2)
	sum2 := sliceInt(slice2).Sum()
	fmt.Println(sum2)

}

type (
	sliceInt []int
	myMap    map[string]any
)

func (si sliceInt) Feed() {
	for i := range si {
		si[i] = rand.Int()
	}
}

func (si sliceInt) Sum() int {
	sum := 0
	for _, v := range si {
		sum += v
	}
	return sum
}

// class A , M1
// Class B : A -> M2
// Class C: A,B --> M3
