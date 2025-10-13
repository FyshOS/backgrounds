package builtin

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
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
	return "Cross"
}

func Dark() fyne.CanvasObject {
	img := canvas.NewImageFromResource(resourceDarkPng)
	img.FillMode = canvas.ImageFillCover
	img.ScaleMode = canvas.ImageScaleFastest
	return img
}

func Light() fyne.CanvasObject {
	img := canvas.NewImageFromResource(resourceLightPng)
	img.FillMode = canvas.ImageFillCover
	img.ScaleMode = canvas.ImageScaleFastest
	return img
}
