package main

import "fmt"

type Vertex2 struct {  // struct is a collection of fields
	X int
	Y int
}

func main() {
	v := Vertex2{1,2}
	v.X = 4
	fmt.Println(v.X) // 4

}

