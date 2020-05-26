package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// Without initializers
var c, python, java bool

// With initializers
var k, l int = 1, 2

var (
	// ToBe is a bool variable
	ToBe bool = false
	// MaxInt is a uint64 variable
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// Pi is a const
const Pi = 3.14

// Numeric constants
const (
	// Big is a large number
	Big = 1 << 100 // a binary number that is 1 followed by 100 zeros
	// Small is a small number
	Small = Big >> 99 // 1 << 1, or 2
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	// Without initializers
	var i int
	fmt.Println(i, c, python, java)
	// 0 false false false

	// With initializers
	var ruby = true
	fmt.Println(k, l, ruby)

	// Short assignment statement
	// Same as var m = 3
	// It can be used inside a function
	m := 3
	fmt.Println(m)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Variable without an explicit initial value are given their zero value
	var i2 int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i2, f, b, s)
	// 0 0 false ""

	// Type conversions
	var n, o int = 3, 4
	var p float64 = math.Sqrt(float64(n*n + o*o))
	var q uint = uint(p)
	fmt.Println(n, o, q)

	// Type inference
	qq := 42
	r := 3.142
	ss := 0.867 + 0.5i
	fmt.Printf("qq is of type %T\n", qq)
	fmt.Printf("r is of type %T\n", r)
	fmt.Printf("ss is of type %T\n", ss)

	// Constants
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	// Numeric constants
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
