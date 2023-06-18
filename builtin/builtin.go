package builtin

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

type Builtin struct{}

func Default() *Builtin {
	return &Builtin{}
}

func (b *Builtin) Load(_ fyne.Theme, v fyne.ThemeVariant) fyne.CanvasObject {
	if v == theme.VariantLight {
		return Light()
	} else {
		return Dark()
	}
}

func (b *Builtin) Name() string {
	return "Hexagons"
}

func Dark() fyne.CanvasObject {
	img := canvas.NewImageFromResource(resourceDarkPng)
	img.FillMode = canvas.ImageFillContain
	img.ScaleMode = canvas.ImageScaleFastest
	return container.NewMax(canvas.NewRectangle(color.NRGBA{R: 0x24, G: 0x24, B: 0x24, A: 0xFF}), img)
}

func Light() fyne.CanvasObject {
	img := canvas.NewImageFromResource(resourceLightPng)
	img.FillMode = canvas.ImageFillContain
	img.ScaleMode = canvas.ImageScaleFastest
	return container.NewMax(canvas.NewRectangle(color.NRGBA{R: 0x1a, G: 0x73, B: 0xe8, A: 0xFF}), img)
}
