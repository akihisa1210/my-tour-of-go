package main

import "fmt"

func deferBasic() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

func deferStacking() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		fmt.Println("defer", i)
		defer fmt.Println(i)
	}
}

func main() {
	deferBasic()
	deferStacking()
}
