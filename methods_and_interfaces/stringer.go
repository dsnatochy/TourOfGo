package main

import "fmt"

// One of the most ubiquitous interfaces is *Stringer* in *fmt* package
// type Stringer interface {
//   String() string
// }
//
// A Stringer is a type that can describe itself as a string. The *fmt*
// package (and many others) look for this interface to print values


type Person struct{
	Name string
	Age int
}

func (p Person) String() string{ //  equivalent of Java's toString() method
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a,z)
}