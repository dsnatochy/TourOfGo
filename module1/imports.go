package module1

import (
	"fmt"
	"math"
)

// this works, but bad style
// import "math"
// import "fmt"

// preferred style
// import (
//    "fmt"
//    "math/rand"
// )

func main() {
	fmt.Printf("Now  you have  %g problems. \n", math.Sqrt(7))
}
