package main

import (
	"fmt"
	"math"
)

// Interface
// https://tour.golang.org/methods/9

// Abser is an interface.
type Abser interface {
	Abs() float64
}

func interfaces() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// a = v

	fmt.Println(a.Abs())
}

// MyFloat is my float!
type MyFloat float64

// Abs for MyFloat.
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Vertex is my vertex!
type Vertex struct {
	X, Y float64
}

// Abs for Vertex.
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Interfaces are implemented implicitly
// https://tour.golang.org/methods/10

// I is an interface.
type I interface {
	M()
}

// T is a struct.
type T struct {
	S string
}

// M prints T's content.
// This method means type T implements the interface I,
// but we don't need to explisitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func interfacesImplementedImplicitly() {
	var i I = T{"Hello"}
	i.M()
}

// Interface values
// https://tour.golang.org/methods/11

// T2 is my struct!
type T2 struct {
	S string
}

// M implemented on T2.
func (t2 *T2) M() {
	fmt.Println(t2.S)
}

// F is my float!
type F float64

// M implemented on F.
func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func interfaceValues() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

// Interface values with nil underlying values
// https://tour.golang.org/methods/12

// T3 is my struct!
type T3 struct {
	S string
}

// M implemented on T3.
func (t3 *T3) M() {
	if t3 == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t3.S)
}

func interfaceValuesWithNil() {
	var i I

	var t3 *T3
	i = t3
	describe(i)
	i.M()

	i = &T3{"Hello"}
	describe(i)
	i.M()
}

// Nil interface values
// https://tour.golang.org/methods/13

// It causes a null pointer error.
// func nilInterfaceValues() {
// 	var i I
// 	describe(i)
// 	i.M()
// }

// The empty interface
// https://tour.golang.org/methods/14

func emptyInterface() {
	var i interface{}
	describe2(i)

	i = 42
	describe2(i)

	i = "hello"
	describe2(i)
}

func describe2(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// Type assertions
// https://tour.golang.org/methods/15

func typeAssertions() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// f = i.(float64) // panic
	// fmt.Println(f)
}

// Type switches
// https://tour.golang.org/methods/16

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func typeSwitches() {
	do(21)
	do("hello")
	do(true)
}

// Stringers
// https://tour.golang.org/methods/17

// Person is a person.
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func stringers() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblexrox", 9001}
	fmt.Println(a, z)
}

func main() {
	interfaces()
	interfacesImplementedImplicitly()
	interfaceValues()
	interfaceValuesWithNil()

	// panic: runtime error: invalid memory address or nil pointer dereference
	// nilInterfaceValues()

	emptyInterface()
	typeAssertions()
	typeSwitches()
	stringers()
}
