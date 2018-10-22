// https://tour.go-zh.org/methods/25

package main

import "golang.org/x/tour/pic"
import "image"
import "image/color"

type Image struct{}

func (img Image) ColorModel() color.Model {
	return color.Alpha16Model
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 200, 100)
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
