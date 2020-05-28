package main

import (
	"golang.org/x/tour/pic"
)

// Pic returns dx * dy size picture.
func Pic(dx, dy int) [][]uint8 {
	row := make([]uint8, dx)
	canvas := make([][]uint8, dy)

	for i := range canvas {
		canvas[i] = row
	}

	for i := range canvas {
		for j := range canvas[i] {
			canvas[i][j] = uint8((i + j) / 2)
			// canvas[i][j] = uint8(i * j)
			// canvas[i][j] = uint8(math.Pow(float64(i), float64(j)))
		}
	}
	return canvas
}

func main() {
	pic.Show(Pic)
}
