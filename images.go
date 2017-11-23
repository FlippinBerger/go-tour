package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type MyImage struct {
	w, h int
}

func main() {
	m := MyImage{100, 200}
	pic.ShowImage(&m)
}

func (img *MyImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)
}

func (img *MyImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (img *MyImage) At(x, y int) color.Color {

	return color.RGBA{uint8(x), uint8(y), 255, 255}
}
