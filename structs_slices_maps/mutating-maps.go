package main

import "fmt"

func main() {
	m := make(map[string]int)

	// insert or update an element in map
	m["Answer"] = 42
	fmt.Println("The value: ", m["Answer"]) // The value:  42

	m["Answer"] = 48
	// retrieve an element
	fmt.Println("The value: ", m["Answer"]) // The value:  48

	// delete an element
	delete(m, "Answer")
	fmt.Println("The value: ", m["Answer"]) // The value:  0

	//m["Answer"] = 0
	//fmt.Println("The value: ", m["Answer"]) // The value:  48

	// test if a key is present with a 2 value assignment
	v, ok := m["Answer"]
	fmt.Println("The value: ", v, "Present?", ok) // The value:  0 Present? false
}
