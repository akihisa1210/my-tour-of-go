package main

import "fmt"

func forBasic() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += sum
	}
	fmt.Println(sum)
}

func forAsWhile() {
	sum := 1
	for sum < 1000 {
		sum += sum
		println(sum)
	}
	fmt.Println(sum)
}

// Infinite loop
//
// func forEver() {
// 	for {
// 	}
// }

func main() {
	forBasic()
	forAsWhile()
	// forEver()
}
