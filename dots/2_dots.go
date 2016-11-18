// See
//    https://blog.golang.org/go-imagedraw-package
// for some more details on this
package main

import (
	"image"
	"image/color"
	"image/draw"
)

func (d *Dots) DrawDots(target draw.Image) {
	dotColor := color.NRGBA{0, 0, 0, 255}
	source := image.NewUniform(dotColor)

	draw.DrawMask(target, target.Bounds(), source, image.ZP, DotMask{d}, image.ZP, draw.Over)
}

type DotMask struct {
	*Dots
}

func (m DotMask) ColorModel() color.Model {
	return color.AlphaModel
}
func (m DotMask) Bounds() image.Rectangle {
	size := (m.Size+1)*BOX_SIZE + m.Size*LINE_SIZE
	return image.Rect(0, 0, size, size)
}
func (m DotMask) At(x, y int) color.Color {
	x = x % (BOX_SIZE + LINE_SIZE)
	y = y % (BOX_SIZE + LINE_SIZE)
	if x > BOX_SIZE && y > BOX_SIZE {
		return color.Alpha{255}
	}
	return color.Alpha{0}
}
