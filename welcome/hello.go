package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, 世界")

	// If you are on the Go Playground, time returns deterministic value.
	fmt.Println("Welcome to the playground!")
	fmt.Println("This time is", time.Now())

}
