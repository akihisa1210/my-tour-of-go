package main

import (
	"fmt"
	"strings"
)

func sliceBasic() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] // primes[1] to primes[3]
	fmt.Println(s)
}

func sliceRefArray() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func sliceLiterals() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func sliceDefaults() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[0:6]
	fmt.Println(s)

	s = s[:6]
	fmt.Println(s)

	s = s[0:]
	fmt.Println(s)

	s = s[:]
	fmt.Println(s)

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

func sliceLenCap() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice to give it zero length
	s = s[:0]
	printSlice(s)

	// Extend its length
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	// Over the capacity
	// panic: runtime error: slice bounds out of range [:10] with capacity 4
	// s = s[:10]
	// printSlice(s)
}

func sliceNil() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func sliceMake() {
	a := make([]int, 5)
	printSlice2("a", a)

	b := make([]int, 0, 5)
	printSlice2("b", b)

	c := b[:2]
	printSlice2("c", c)

	d := c[2:5]
	printSlice2("d", d)
}

func sliceSlice() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "0"
	board[1][2] = "X"
	board[1][0] = "0"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], "  "))
	}
}

func sliceAppending() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)

	// backing array is small
	var a = [4]int{1, 2, 3, 4}
	var s2 = a[:]
	printSlice(s2)
	s2 = append(s2, 5, 6, 7, 8)
	printSlice(s2)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func main() {
	sliceBasic()
	sliceRefArray()
	sliceLiterals()
	sliceDefaults()
	sliceLenCap()
	sliceNil()
	sliceMake()
	sliceSlice()
	sliceAppending()
}
