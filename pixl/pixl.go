package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"thegadri.io/pixl/apptype"
	"thegadri.io/pixl/pxcanvas"
	"thegadri.io/pixl/swatch"
	"thegadri.io/pixl/ui"
)

func main() {
	pixlApp := app.New()
	pixlWindow := pixlApp.NewWindow("pixl")

	state := apptype.State{
		BrushColor:     color.NRGBA{255, 255, 255, 255},
		SwatchSelected: 0,
	}

	pixlCanvasConfig := apptype.PxCanvasConfig{
		DrawingArea:  fyne.NewSize(600, 600),
		CanvasOffset: fyne.NewPos(0, 0),
		PxRow:        10,
		PxCols:       10,
		PxSize:       30,
	}

	pixlCanvas := pxcanvas.NewPxCanvas(&state, pixlCanvasConfig)

	appInit := ui.AppInit{
		PixelCanvas: pixlCanvas,
		PixlWindow:  pixlWindow,
		State:       &state,
		Swatches:    make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)

	appInit.PixlWindow.ShowAndRun()
}
