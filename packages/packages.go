package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("My favorite nubmer is", rand.Intn(10))

	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	// A name is exported if it begins with a capital latter.
	// Error: cannot refer to unexported name math.pi
	// fmt.Println(math.pi)
	fmt.Println(math.Pi)
}
