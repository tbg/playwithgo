package main

import (
	"image"
	"image/color"

	"code.google.com/p/go-tour/pic"
)

type Image struct {
	w int
	h int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x * y), uint8(x * y), 255, 255}
}

func main() {
	m := Image{1000, 1000}
	pic.ShowImage(m)
}
