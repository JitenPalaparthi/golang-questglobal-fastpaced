package main

import "fmt"

func main() {

	r1 := Rect{L: 100.11, B: 102.45}
	a1 := r1.Area()
	p1 := r1.Perimeter()
	fmt.Println("Area:", a1)
	fmt.Println("Perimeter:", p1)
	fmt.Println("Area in r1:", r1.A)
	fmt.Println("Perimeter in r1:", r1.P)

}

type Rect struct {
	L, B, A, P float32
}

func (r *Rect) Area() float32 {
	r.A = r.L * r.B
	return r.A
}

func (r *Rect) Perimeter() float32 {
	r.P = 2 * (r.L * r.B)
	return r.P
}
