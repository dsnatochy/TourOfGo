package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	var prevValue float64

	for i:=0; i<10; i++{
		z -= (z*z - x) / (2*z)
		if z == prevValue{
			break
		}
		prevValue = z
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(9))

}