package main

import (
	"fmt"
	"reflect"
)

func main() {

	var num1 integer = 100

	fmt.Println(num1)

	var a int = 10

	var b integer = 20

	c := sum(a, b)
	fmt.Println(c)

	var mi1 myint = 5000

	str1 := mi1.ToString()

	fmt.Println(str1)

	str2 := myint(num1).ToString()
	fmt.Println(str2, "type:", reflect.TypeOf(str2))

}

type integer = int // alias

func sum(a, b integer) integer {
	return a + b
}

// type any = interface{}

// type rune = int32

// type byte = uint8

type myint int

func (mi myint) ToString() string {
	return fmt.Sprint(mi)
}
