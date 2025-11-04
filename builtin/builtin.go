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
	var img *canvas.Image
	if v == theme.VariantLight {
		img = canvas.NewImageFromResource(resourceLightPng)
	} else {
		img = canvas.NewImageFromResource(resourceDarkPng)
	}

	img.FillMode = canvas.ImageFillCover
	img.ScaleMode = canvas.ImageScaleFastest
	return img
}

func (b *Builtin) Name() string {
	return "Cross"
}
