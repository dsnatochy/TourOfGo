package module1

import (
	"math"
	"fmt"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y + 1))
	z := uint(f)  // conversion
	fmt.Println(x, y, z, f)
}
