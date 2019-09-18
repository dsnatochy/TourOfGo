package main

import (
	"math"
	"fmt"
)

func compute(fn func(float64, float64) float64) float64{
	return fn(3,4)
}

func main() {
	// functions are values too
	hypot := func(x,y float64) float64{
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5,12)) // 13

	// functions can be passed around just like other values
	fmt.Println(compute(hypot)) // 5
	fmt.Println(compute(math.Pow)) // 81

}