package builtin

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

// Builtin type represents a background with the default design, it has light and dark variants

type Builtin struct{}

// Default returns the default builtin background
func Default() *Builtin {
	return &Builtin{}
}

// Load returns a background image ojbect set up to scale appropriately
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

// Name returns a string representation of this design
func (b *Builtin) Name() string {
	return "Cross"
}
