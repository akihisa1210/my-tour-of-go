package main

import "golang.org/x/tour/reader"

// MyReader is a type for implementing my Read()
type MyReader struct{}

func (r MyReader) Read(n []byte) (int, error) {
	for i := range n {
		n[i] = 0x41
	}
	return len(n), nil
}

func main() {
	reader.Validate(MyReader{})
}
