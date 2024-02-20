package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"zerotomastery.io/pixl/swatch"
)

func BuildSwatches(app *AppInit) *fyne.Container {
	canvasSwatches := make([]fyne.CanvasObject, 0, 64)
	for i := 0; i < cap(app.Swatches); i++ {
		initialColor := color.NRGBA{255, 255, 255, 255}
		s := swatch.NewSwatch(app.State, initialColor, i, func(s *swatch.Swatch) {
			for j := 0; j < len(app.Swatches); j++ {
				app.Swatches[j].Selected = false
				canvasSwatches[j].Refresh()
			}
			app.State.SwatchSelected = s.SwatchIndex
			app.State.BrushColor = s.Color
		})
		if i == 0 {
			s.Selected = true
			app.State.SwatchSelected = 0
			s.Refresh()
		}
		app.Swatches = append(app.Swatches, s)
		canvasSwatches = append(canvasSwatches, s)
	}

	return container.NewGridWrap(fyne.NewSize(20, 20), canvasSwatches...)
}
