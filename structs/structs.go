package main

import "fmt"

// Vertex is a point where two lines meet to form an angle.
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{3, 4}
	v.X = 5
	fmt.Println(v.X)

	p := &v
	p.X = 1e9
	fmt.Println(v)

	// struct literals
	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		pp = &Vertex{1, 2} // has type *Vertex
	)

	fmt.Println(v1, pp, v2, v3)
}
