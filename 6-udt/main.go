package main

import "fmt"

func main() {
	p1 := new(Person)
	//p1 := Person{}
	//p1 := &Person{}

	p1.Id = "101"
	p1.Name = "Jiten"
	p1.Email = "Jitenp@outlook.com"
	p1.Address.City = "Bangalore"
	p1.Address.Country = "India"
	p1.Address.Line1 = "Road-1,Yeshvantpur"
	p1.Address.PinCode = "560086"
	p1.Address.Status = "active"
	p1.Status = "okay"

	e1 := new(Employee)
	//p1 := Person{}
	//p1 := &Person{}

	e1.Id = "101"
	e1.Name = "Jiten"
	e1.Email = "Jitenp@outlook.com"
	e1.City = "Bangalore"
	e1.Country = "India"
	e1.Line1 = "Road-1,Yeshvantpur"
	e1.PinCode = "560086"
	e1.Status = "active"
	e1.Address.Status = "active"

	t1 := TypeMaster{int: 101, string: "main-type", integer: 12123123, string2: "sub-type"}
	fmt.Println(t1)

	e2 := struct {
		Id   int
		Name string
	}{
		Id:   101,
		Name: "Jiten",
	}

	fmt.Println(e2)

}

type Person struct {
	Id, Name, Email, Status string
	Address                 Address
}

type Address struct {
	Line1, City, Country, PinCode, Status string
}

type Employee struct {
	Id, Name, Email, Status string
	Address                 // promoted
}

type TypeMaster struct {
	int
	string
	integer // alias
	string2 // alias
}

type integer = int // alias

type string2 string // new types

type A struct {
	F1 integer
	F2 string
	F3 float32
	B  struct {
		F1 string
		F2 bool
		F3 int
	}
}

// type A struct {
// 	gorm.Model
// 	Name  string
// 	Email string
// }

// class A ID,Name
// class B extends A : Line1

// b= new B() ; b.ID, b.Name
