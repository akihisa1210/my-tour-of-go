package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13r *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = r13r.r.Read(b)

	for i, v := range b {
		switch {
		case (v >= 0x41 && v <= 0x4d) || (v >= 0x61 && v <= 0x6d):
			b[i] += 13
		case (v >= 0x4e && v <= 0x5a) || (v >= 0x6e && v <= 0x7a):
			b[i] -= 13
		default:
			continue
		}
	}
	return len(b), err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	_, err := io.Copy(os.Stdout, &r)
	// Result:
	// You cracked the code!Lbh penpxrq gur pbqr!
	if err != nil {
		println(err)
	}
}
