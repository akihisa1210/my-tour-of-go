package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount is counting word in the given string.
func WordCount(s string) map[string]int {
	words := strings.Fields(s)

	result := make(map[string]int)

	for _, v := range words {
		_, exist := result[v]
		if exist {
			result[v]++
			continue
		}
		result[v] = 1
	}

	return result
}

func main() {
	WordCount("I am learning Go!")
	wc.Test(WordCount)
}
