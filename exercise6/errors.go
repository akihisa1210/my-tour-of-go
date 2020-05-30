package main

import "fmt"

// Sqrt is my sqrt function.
// It should return a non-nil error value when given a negative number
// because it doesn't support complex numbers.
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	oldz := 0.0
	for i := 0; i < 10 && z != oldz; i++ {
		oldz = z
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

// ErrNegativeSqrt is an error Sqrting negative number.
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// // Sqrt is a function that if a number x is given
// // it returns z for which z^2 is most nearly x.
// func Sqrt(x, z float64) float64 {
// 	// fmt.Printf("Sqrt(%f, %f) start\n", x, z)
// 	oldz := 0.0
// 	// i < 100 が成り立っている間はループする
// 	// z == oldz が成り立ったらループを抜ける
// 	for i := 0; i < 10 && z != oldz; i++ {
// 		oldz = z
// 		z -= (z*z - x) / (2 * z)
// 		// fmt.Printf(" %d ", i)
// 	}
// 	// fmt.Println("finished!")
// 	return z
// }

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
