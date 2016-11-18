package main

import (
	//"image"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"image/color"
	"image/draw"
)

func (d *Dots) DrawInfo(img draw.Image) {
	draw2d.SetFontFolder(".")
	gc := draw2dimg.NewGraphicContext(img)
	gc.SetFontData(draw2d.FontData{
		Name:   "DroidSansMono",
		Family: draw2d.FontFamilyMono,
	})
	gc.SetFillColor(color.RGBA{0, 0, 0, 255})
	gc.SetFontSize(15)
	var str string
	if d.P1Turn {
		str = "Player One's Turn"
	} else {
		str = "Player Two's Turn"
	}
	gc.FillStringAt(str, 5, 20)
}
