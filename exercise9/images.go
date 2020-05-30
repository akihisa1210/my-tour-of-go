package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Image is my image type.
type Image struct{}

// ColorModel for my Image type.
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds for my Image type.
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)

}

// At for my Image type.
func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}