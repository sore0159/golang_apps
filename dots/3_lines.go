package main

import (
	//"image"
	"github.com/llgcode/draw2d/draw2dimg"
	"image/color"
	"image/draw"
)

func (d *Dots) DrawBoxes(img draw.Image) {
	gc := draw2dimg.NewGraphicContext(img)
	for _, bx := range d.Boxes {
		switch bx[2] {
		case 1:
			gc.SetFillColor(color.NRGBA{205, 150, 150, 255})
		case 2:
			gc.SetFillColor(color.NRGBA{150, 205, 150, 255})
		default:
			continue
		}
		x1, y1 := PtLoc(bx[0], bx[1])
		x2, y2 := PtLoc(bx[0]+1, bx[1]+1)
		gc.MoveTo(x1, y1)
		gc.LineTo(x1, y2)
		gc.LineTo(x2, y2)
		gc.LineTo(x2, y1)
		gc.Close()
		gc.Fill()
	}
}

func (d *Dots) DrawLines(img draw.Image) {
	gc := draw2dimg.NewGraphicContext(img)
	gc.SetLineWidth(LINE_SIZE)
	for _, ln := range d.Lines {
		switch ln[2] {
		case 1:
			gc.SetStrokeColor(color.NRGBA{255, 0, 0, 255})
		case 2:
			gc.SetStrokeColor(color.NRGBA{0, 255, 0, 255})
		default:
			continue
		}
		x1, y1, x2, y2 := d.LineDots(ln)
		x, y := PtLoc(x1, y1)
		gc.MoveTo(x, y)
		x, y = PtLoc(x2, y2)
		gc.LineTo(x, y)
		gc.Stroke()
	}
}
