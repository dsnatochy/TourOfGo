package main

import (
	"math"
	"fmt"
)

type Vertex2 struct {
	X, Y float64
}

// Remember: a method is just a function with a receiver argument.
//Here's Abs written as a regular function with no change in functionality.
func Abs(v Vertex2) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex2{3, 4}
	fmt.Println(Abs(v))
}
