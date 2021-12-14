package main

import (
	"fmt"
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

func (m Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 255, 255)
}

func (m Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (m Image) At(x, y int) color.Color {
	c := uint8(x + y)
	return color.RGBA{c, c, 255, 255}
}

func main() {

	m1 := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m1.Bounds())
	fmt.Println(m1.At(0, 0).RGBA())
	m := Image{}
	pic.ShowImage(m)
}