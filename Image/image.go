package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{
	Height, Weight int
}

func (p Image) Bounds () image.Rectangle {
	return image.Rect(0, 0, p.Weight, p.Height)
}

func (p Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (p Image) At(x, y int) color.Color {
	c := uint8(x ^ y)
	return color.RGBA{c, c, 255, 255}
}

func main() {
	m := Image{100, 100}
	pic.ShowImage(m)
}