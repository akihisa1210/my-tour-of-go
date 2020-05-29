package main

import "fmt"

func fibonacci() func() int {
	count := 0
	a := 0
	b := 0
	result := 0
	return func() int {
		count++
		switch {
		case count == 1:
			a = 0
			return a
		case count == 2:
			b = 1
			return b
		default:
			result = a + b
			a = b
			b = result
			return result
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
