package main

import "fmt"

// A type assertion provides access to an interface value's underlying concrete value.
//
//t := i.(T)
//This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t.
//
//If i does not hold a T, the statement will trigger a panic.

func main() {
	var i interface{} = "hello"

	s:= i.(string)
	// s:= i.(int) // wrong type = panic: interface conversion: interface {} is string, not int
	fmt.Println(s) // hello

	s, ok := i.(string)  // kind of like "instance of"
	if ok {
		fmt.Println(s,ok) // hello true
	}

	f, ok := i.(float64) // equivalent to: Float f = 0; boolean ok = false; if (i instanceof Float){ f = i; ok = true}
	fmt.Println(f, ok) // 0, false

	var x testI
	x = &testType{"testType value"}
	x.testMethod() // from testMethod

	// fmt.Println(x.s) // x.s undefined (type testI has no field or method s)

	z, ok := x.(*testType)
	fmt.Println(z, ok) // &{testType value} true
	fmt.Println(z.s)  // testType value
}

type testI interface {
	testMethod()
}

type testType struct{
	s string
}

func (t *testType) testMethod() {
	fmt.Println("from testMethod")
}