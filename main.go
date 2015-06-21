package main

import (
	"image"
	"image/color"
	"log"
	"math"

	"github.com/jdf/filmore"
	"github.com/llgcode/draw2d"
)

const (
	fontSize = 64
	width    = 700
	height   = 380

	fontPath = "BebasNeue Bold.ttf"
	//fontPath = "alpha_echo.ttf"
)

var (
	text = []string{"Absinthe", "Makes", "The heart", "Grow flounder"}
)

func main() {
	font, err := filmore.NewFontFromFile(fontPath, fontSize)
	if err != nil {
		log.Fatal(err)
	}
	// Initialize the graphic context on an RGBA image
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	gc := draw2d.NewGraphicContext(img)

	// Draw a rounded rectangle using default colors
	draw2d.RoundRect(gc, 5, 5, width-5, height-5, 10, 10)
	gc.FillStroke()

	gc.SetStrokeColor(image.Black)
	gc.SetFillColor(color.RGBA{0, 255, 0, 255})

	// Warp function
	f := func(x, y float64) (nx, ny float64) {
		xMag := x / 30
		mid := height / 2.0
		yMag := 1.25 * math.Abs(y-mid) / mid
		nx = x + xMag*yMag*math.Cos(x/25.0)
		ny = y + xMag*yMag*math.Sin(x/20.0)
		return
	}

	doPath := func(p filmore.TextPath) {
		for _, op := range p.PathOps {
			tx, ty := f(op.X(), op.Y())
			switch op := op.(type) {
			case filmore.MoveTo:
				gc.MoveTo(tx, ty)
			case filmore.LineTo:
				gc.LineTo(tx, ty)
			case filmore.QuadCurveTo:
				tcx, tcy := f(op.ControlX(), op.ControlY())
				gc.QuadCurveTo(tx, ty, tcx, tcy)
			}
		}
	}

	spacing := fontSize * 1.3
	y := 12 + spacing
	for _, s := range text {
		p := font.CreateTextPath(s, 20, y)
		doPath(p)
		gc.Fill()
		doPath(p)
		gc.Stroke()
		y += spacing
	}
	// Save to png
	draw2d.SaveToPngFile("helloworld.png", img)
}
