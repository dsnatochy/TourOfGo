package main

import "fmt"

type Vertex3 struct {  // struct is a collection of fields
	X int
	Y int
}

func main() {
	v := Vertex3{1, 2}

	p:= &v // pointer to the struct

	p.X = 1e9 // struct can be access via pointer (1e9 = Math.pow(10,9))
	fmt.Println(p) //{1000000000 2}

}
