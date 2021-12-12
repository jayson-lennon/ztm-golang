package ui

import (
	"fyne.io/fyne/v2"
	"zerotomastery.io/pixl/apptype"
	"zerotomastery.io/pixl/pxcanvas"
	"zerotomastery.io/pixl/swatch"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
