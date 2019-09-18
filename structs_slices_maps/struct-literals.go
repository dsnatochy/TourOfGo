package main

import "fmt"

type Vertex4 struct{
	X , Y int
}

var(
	v1 = Vertex4{1, 2}
	v2 = Vertex4{Y:1, X:2} // order is irrelevant when "Name: " syntax is used
	v3 = Vertex4{}         // X:0, Y:0
	p = &Vertex4{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3) //{1 2} &{1 2} {2 1} {0 0}
}