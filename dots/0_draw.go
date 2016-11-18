package main

import (
	"errors"
	"image"
	"image/png"
	"os"
)

func (d *Dots) Draw() error {
	img, err := d.CreateImg()
	if err != nil {
		return err
	}
	file, err := os.Create("dots.png")
	if err != nil {
		return err
	}
	defer file.Close()
	return png.Encode(file, img)
}

const (
	BOX_SIZE  = 50
	LINE_SIZE = 5
)

func (d *Dots) CreateImg() (image.Image, error) {
	if !d.Valid() {
		return nil, errors.New("Invalid dots data for drawing")
	}
	size := (d.Size+1)*BOX_SIZE + d.Size*LINE_SIZE

	rect := image.Rect(0, 0, size, size)
	img := image.NewRGBA(rect)
	//               //
	d.DrawBackground(img)
	d.DrawInfo(img)
	d.DrawBoxes(img)
	d.DrawLines(img)
	d.DrawDots(img)
	return img, nil
}

func PtLoc(x, y int) (float64, float64) {
	return float64((BOX_SIZE+LINE_SIZE)*x - LINE_SIZE/2),
		float64((BOX_SIZE+LINE_SIZE)*y - LINE_SIZE/2)
}
