package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{
	Height, Width int
}

//func (image *Image) ShowImage() {
//}

func (i Image) Bounds() image.Rectangle{
	return image.Rect(0,0, i.Height,i.Width)
}

func (image Image) ColorModel() color.Model{
	return color.RGBAModel
}

func (image Image) At(x, y int) color.Color{
	c := uint8(x ^ y)
	return color.RGBA{c, c, 255, 255}
}

func main() {
	m := Image{256,256}
	pic.ShowImage(m)
}
