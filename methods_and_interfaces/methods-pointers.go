package main

import (
	"math"
	"fmt"
)

type Vertex3 struct {
	X, Y float64
}

func (v Vertex3) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Methods with pointer receivers can modify the value to which the receiver points (as Scale does here).
// Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
//
//Try removing the * from the declaration of the Scale function on line 16 and observe how the program's behavior changes.
//
//With a value receiver, the Scale method operates on a copy of the original Vertex value.
// (This is the same behavior as for any other function argument.)
// The Scale method must have a pointer receiver to change the Vertex value declared in the main function.
func (v *Vertex3) Scale(f float64)  {
	v.X = v.X * f
	v.Y *= f
}

func main() {
	v := Vertex3{3, 4}
	v.Scale(10) // Implicitly converted to (&v).Scale(10)

	fmt.Println(v.Abs()) //50
}