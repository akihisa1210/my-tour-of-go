package main

import "fmt"

func mapBasic() {
	// Vertex is an address.
	type Vertex struct {
		Lat, Long float64
	}

	var m map[string]Vertex

	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

func mapLiterals() {
	// Vertex is an address.
	type Vertex struct {
		Lat, Long float64
	}

	var m = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)
}

func mapLiteralsContinued() {
	// Vertex is an address.
	type Vertex struct {
		Lat, Long float64
	}

	var m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m)
}

func mapMutating() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	m["Answer"] = 0
	fmt.Println("The value:", m["Answer"])

	v2, ok := m["Answer"]
	fmt.Println("The value:", v2, "Present?", ok)
}

func main() {
	mapBasic()
	mapLiterals()
	mapLiteralsContinued()
	mapMutating()
}
