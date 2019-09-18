package main

import (
	"math"
	"fmt"
)

// An interface type is defined as a set of method signatures.

// A value of interface type can hold any value that implements those methods.

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat2(-math.Sqrt2)
	v := Vertex5{3, 4}

	a = f // a MyFloat2 implements Abser
	a = &v // a *Vertex5 implements Abser

	// In the following line, v is a Vertex5 (not *Vertex5)
	// and does NOT implement Abser
	a = v

	fmt.Println(a.Abs())
}

type MyFloat2 float64

type Vertex5 struct {
	X, Y float64
}

func (f MyFloat2) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex5) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}