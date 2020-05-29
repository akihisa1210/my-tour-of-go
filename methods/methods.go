package main

import (
	"fmt"
	"math"
)

// Vertex is a pair of float64.
type Vertex struct {
	X, Y float64
}

// Abs returns the length from origin to the vertex.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// AbsFunc is the method which has same functionalry with Abs()
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func methods() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))
}

// MyFloat is my float!
type MyFloat float64

// Abs is for MyFloat.
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func methodsContinued() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

// Scale multiplies Vertex.
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// ScaleFunc is function styled Scale().
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func pointerReceivers() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

func pointerReceiversFunctions() {
	v := Vertex{3, 4}
	ScaleFunc(&v, 10)
	fmt.Println(AbsFunc(v))
}

func methodsAndPointerIndirection() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}

func methodsAndPointerIndirection2() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}

func choosingValueOrPointerReceiver() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}

func main() {
	methods()
	methodsContinued()
	pointerReceivers()
	pointerReceiversFunctions()
	methodsAndPointerIndirection()
	methodsAndPointerIndirection2()
	choosingValueOrPointerReceiver()
}
