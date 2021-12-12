package util

import (
	"image"
	"image/color"
)

func GetImageColors(img image.Image) map[color.Color]struct{} {
	colors := make(map[color.Color]struct{})
	var empty struct{}
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			colors[img.At(x, y)] = empty
		}
	}
	return colors
}
