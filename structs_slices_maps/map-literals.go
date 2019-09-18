package main


import "fmt"

type Vertex6 struct {
	Lat, Long float64
}


func main() {
	// Map literals are like struct literals, but the keys are required.
	var m = map[string]Vertex6{
		"Bell Labs": Vertex6{
			40.68433, -74.39967,
		},
		"Google": Vertex6 {
			37.42202, -122.08408,
		},
	}
	fmt.Println(m) // map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]

	// If the top-level type is just a type name, you can omit it from the elements of the literal.
	var n = map[string]Vertex6{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(n)
}

