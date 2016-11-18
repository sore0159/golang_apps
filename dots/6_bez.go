package main

import (
	"fmt"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"image/color"
	"image/png"
	"io"
	"math"
)

func DrawBez(w io.Writer, pts [][2]float64) {
	rect := image.Rect(0, 0, 500, 500)
	img := image.NewRGBA(rect)
	gc := draw2dimg.NewGraphicContext(img)
	if len(pts) == 3 {
		gc.SetLineWidth(2)
		gc.SetStrokeColor(color.NRGBA{130, 130, 130, 255})
		gc.MoveTo(pts[0][0], pts[0][1])
		gc.LineTo(pts[1][0], pts[1][1])
		gc.LineTo(pts[2][0], pts[2][1])
		gc.Stroke()
		//
		gc.SetStrokeColor(color.Black)
		gc.SetLineWidth(5)
		gc.MoveTo(pts[0][0], pts[0][1])
		gc.QuadCurveTo(pts[1][0], pts[1][1], pts[2][0], pts[2][1])
		gc.Stroke()
		//
		gc.SetLineWidth(1)
		gc.SetFillColor(color.NRGBA{255, 0, 0, 255})
		Blip(pts[0], gc)
		Blip(pts[2], gc)
		gc.SetFillColor(color.NRGBA{0, 0, 255, 255})
		Blip(pts[1], gc)
	} else if len(pts) == 4 {
		gc.SetLineWidth(2)
		gc.SetStrokeColor(color.NRGBA{130, 130, 130, 255})
		gc.MoveTo(pts[0][0], pts[0][1])
		gc.LineTo(pts[1][0], pts[1][1])
		gc.MoveTo(pts[3][0], pts[3][1])
		gc.LineTo(pts[2][0], pts[2][1])
		gc.Stroke()
		//
		gc.SetStrokeColor(color.Black)
		gc.SetLineWidth(5)
		gc.MoveTo(pts[0][0], pts[0][1])
		gc.CubicCurveTo(pts[1][0], pts[1][1], pts[2][0], pts[2][1], pts[3][0], pts[3][1])
		gc.Stroke()
		//
		gc.SetLineWidth(1)
		gc.SetFillColor(color.NRGBA{255, 0, 0, 255})
		Blip(pts[0], gc)
		Blip(pts[3], gc)
		gc.SetFillColor(color.NRGBA{0, 0, 255, 255})
		Blip(pts[1], gc)
		Blip(pts[2], gc)
	}
	if err := png.Encode(w, img); err != nil {
		fmt.Println("Bez file encode error:", err)
	}
}

func Blip(pt [2]float64, gc draw2d.GraphicContext) {
	r := 8.
	gc.MoveTo(pt[0]+r, pt[1])
	gc.ArcTo(pt[0], pt[1], r, r, 0, 2*math.Pi)
	gc.LineTo(pt[0]+r, pt[1])
	gc.Close()
	gc.FillStroke()
}
