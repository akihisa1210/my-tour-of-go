package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// type Tree struct {
// 	Left *Tree
// 	Value int
// 	Right *Tree
// }

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	if len(ch) < cap(ch) {
		// fmt.Printf("len(ch): %v; cap(ch): %v; t.Value: %v\n", len(ch), cap(ch), t.Value)
		ch <- t.Value
	}
	if t.Right != nil {
		Walk(t.Right, ch)
	}
	return
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	Walk(t1, ch1)
	Walk(t2, ch2)

	for i := 0; i < 10; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	// Test Walk()
	ch := make(chan int, 11)
	go Walk(tree.New(1), ch)
	// This for loop causes a dead lock.
	// for i := range ch {
	// 	fmt.Println(i)
	// }
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// Test Same()
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
