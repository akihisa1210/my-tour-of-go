package main

import "fmt"

// Sqrt is a function that if a number x is given
// it returns z for which z^2 is most nearly x.
func Sqrt(x, z float64) float64 {
	// fmt.Printf("Sqrt(%f, %f) start\n", x, z)
	oldz := 0.0
	// i < 100 が成り立っている間はループする
	// z == oldz が成り立ったらループを抜ける
	for i := 0; i < 10 && z != oldz; i++ {
		oldz = z
		z -= (z*z - x) / (2 * z)
		// fmt.Printf(" %d ", i)
	}
	// fmt.Println("finished!")
	return z
}

func main() {
	// z = 1
	fmt.Println("Sqrt(0, 1): ", Sqrt(0, 1))
	fmt.Println("Sqrt(1, 1): ", Sqrt(1, 1))
	fmt.Println("Sqrt(2, 1): ", Sqrt(2, 1))
	fmt.Println("Sqrt(3, 1): ", Sqrt(3, 1))
	fmt.Println("Sqrt(4, 1): ", Sqrt(4, 1))
	fmt.Println("Sqrt(5, 1): ", Sqrt(5, 1))
	fmt.Println("Sqrt(17, 1): ", Sqrt(17, 1))
	fmt.Println("Sqrt(42, 1): ", Sqrt(42, 1))

	// z = x
	fmt.Println("Sqrt(0, 1): ", Sqrt(0, 0))
	fmt.Println("Sqrt(1, 1): ", Sqrt(1, 1))
	fmt.Println("Sqrt(2, 2): ", Sqrt(2, 2))
	fmt.Println("Sqrt(3, 3): ", Sqrt(3, 3))
	fmt.Println("Sqrt(4, 4): ", Sqrt(4, 4))
	fmt.Println("Sqrt(5, 5): ", Sqrt(5, 5))
	fmt.Println("Sqrt(17, 17): ", Sqrt(17, 17))
	fmt.Println("Sqrt(42, 42): ", Sqrt(42, 42))

	// z = x / 2
	fmt.Println("Sqrt(1, 0.5): ", Sqrt(0, 0))
	fmt.Println("Sqrt(1, 0.5): ", Sqrt(1, 0.5))
	fmt.Println("Sqrt(2, 1): ", Sqrt(2, 1))
	fmt.Println("Sqrt(3, 1.5): ", Sqrt(3, 1.5))
	fmt.Println("Sqrt(4, 2): ", Sqrt(4, 2))
	fmt.Println("Sqrt(5, 2.5): ", Sqrt(5, 2.5))
	fmt.Println("Sqrt(17, 8.5): ", Sqrt(17, 8.5))
	fmt.Println("Sqrt(42, 21): ", Sqrt(42, 21))
}
