package resize

import (
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	//	"image/draw"
	//	"image/png"
	"log"
	//	"os"
	"testing"
)

func TestOne(t *testing.T) {
	log.Println("TEST ONE")
}

func TestTwo(t *testing.T) {
	source, err := draw2dimg.LoadFromPngFile("dots.png")
	if err != nil {
		log.Println("DECODE ERROR ", err)
		return
	}

	sx, sy := 0.25, 0.25
	bounds := source.Bounds()
	img := image.NewRGBA(bounds)

	gc := draw2dimg.NewGraphicContext(img)
	gc.Scale(sx, sy)
	gc.DrawImage(source)

	if err = draw2dimg.SaveToPngFile("dots2.png", img); err != nil {
		log.Println("ENCODE ERROR", err)
		return
	}

}
