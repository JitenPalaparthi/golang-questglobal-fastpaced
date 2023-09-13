package main

import "fmt"

func main() {

	var (
		num1 int     = 100
		num2 myint1  = 101
		num3 myint2  = 102
		num4 float64 = 1232.343
	)

	// int
	str1 := myint1(num1).ToString()
	bytes1 := myint2(num1).ToBytes()
	fmt.Println("To String:", str1)
	fmt.Println("To Bytes:", bytes1)

	str2 := num2.ToString()
	bytes2 := myint2(num2).ToBytes()

	fmt.Println("To String:", str2)
	fmt.Println("To Bytes:", bytes2)

	str3 := myint1(num3).ToString()
	bytes3 := num3.ToBytes()

	fmt.Println("To String:", str3)
	fmt.Println("To Bytes:", bytes3)

	str4 := myint1(num4).ToString()
	bytes4 := myint2(num4).ToBytes()

	fmt.Println("To String:", str4)
	fmt.Println("To Bytes:", bytes4)
}

type (
	myint1 int
	myint2 myint1
)

func (mi1 myint1) ToString() string {
	return fmt.Sprint(mi1)
}

func (mi2 myint2) ToBytes() []byte {
	return []byte(fmt.Sprint(mi2))
}
