package main

import "fmt"

// If the concrete value inside the interface itself is il, the method will
// be called with a nil receiver

// In some languages this would trigger a null pointer exception, but in
// Go it is common to write methods that gracefully handle being called with a nil receiver
//   (as with the method M in this example)
// NOte that an interace value that holds a nil concrete value itself non-nil

type I3 interface {
	M3()
}

type T3 struct {
	S string
}

func (t *T3) M3(){
	if t == nil{
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}


func main() {
	var i I3
	var t *T3
	i = t
	describe2(i)
	i.M3()   // in Java this would be a NullPointerException
	// (<nil>, *main.T3)
	// <nil>

	i = &T3{"Hello"}
	describe2(i)
	i.M3()
	// (&{Hello}, *main.T3)
	// Hello

}

func describe2(i I3){
	fmt.Printf("(%v, %T)\n", i, i)
}