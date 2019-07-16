/*
	Notes 2.25

	Exercise: https://tour.golang.org/methods/25

	Got help from: https://gist.github.com/flc/6437570
*/

package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	width, height int
	colorB, colorA uint8
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), i.colorB, i.colorA}
}

func main() {
	m := Image{111, 111, 255, 255}
	pic.ShowImage(m)
}