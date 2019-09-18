package module1

import "fmt"

var i, j int = 1,2 // A var declaration can include initializers, one per variable.


func main() {
	var c, python, java  = true, false, "no" // If an initializer is present, the type can be omitted
	fmt.Println(i, j, c, python, java)
}
