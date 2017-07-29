package main

import "golang.org/x/tour/pic"
import "image"
import "image/color"

type Image struct{}

func (i Image) Bounds () image.Rectangle {
	return image.Rect(0, 0, 10, 10)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{255, 0, 0, 255} // 固定でいいのかな?
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
