package main

import (
	"math"
	"fmt"
)

type Vertex4 struct {
	X, Y float64
}

// value receiver
func Abs2(v Vertex4) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// pointer receiver
func Scale(v *Vertex4, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// In general all methods on a given type should be either value or pointer receivers

// There are two reasons to use a pointer receiver:
// The first is so that the method can modify the value that its receiver points to.
// The second is to avoid copying the value on each method call.
//   This can be more efficient if the receiver is a large struct, for example.

func main() {
	v := Vertex4{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs2(v)) // 50
}