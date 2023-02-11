//go:generate fyne bundle -package builtin -o bundled.go dark.png
//go:generate fyne bundle -package builtin -o bundled.go -a light.png

package builtin

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

type Builtin struct {}

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
	return canvas.NewImageFromResource(resourceDarkPng)
}

func Light() fyne.CanvasObject {
	return canvas.NewImageFromResource(resourceLightPng)
}
