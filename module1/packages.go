package module1

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	fmt.Println("My favorite number is ", rand.Intn(10)) // will print the same deterministic number

	rand.Seed(time.Now().Unix()) //if the seed is constant, the result will be the same.
	fmt.Println("My favorite number is ", rand.Intn(1000))

}
