package module1

import (
	"fmt"
	"math"
)

func main() {
	//println(math.pi) //fails because "pi" is unexported and can only be referred to from inside "math" package
	fmt.Println(math.Pi)
}
