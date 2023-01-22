package ui

import (
	"fyne.io/fyne/v2"
	"thegadri.io/pixl/apptype"
	"thegadri.io/pixl/swatch"
)

type AppInit struct {
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
