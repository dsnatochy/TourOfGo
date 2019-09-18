package main

import "fmt"

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		// var myError ErrNegativeSqrt = ErrNegativeSqrt(x)
		return x, &ErrNegativeSqrt{x}
	}
	z := 1.0
	var prevValue float64

	for i:=0; i<10; i++{
		z -= (z*z - x) / (2*z)
		if z == prevValue{
			break
		}
		prevValue = z
		//fmt.Println(z)
	}
	return z, nil
}

type ErrNegativeSqrt struct {
	val float64
}

func (e *ErrNegativeSqrt) Error() string{
	return fmt.Sprintf("cannot Sqrt negative number: %v", e.val)
}

func main() {
	fmt.Println(Sqrt(2)) // 1.414213562373095 <nil>
	fmt.Println(Sqrt(-3)) // -3 cannot Sqrt negative number: -3
}