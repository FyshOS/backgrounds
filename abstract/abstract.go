package abstract

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

// Abstract type represents a background with an abstract design, it has light and dark variants
type Abstract struct {
	name string
}

// Load returns a background image ojbect set up to scale appropriately
func (a *Abstract) Load(_ fyne.Theme, v fyne.ThemeVariant) fyne.CanvasObject {
	var img *canvas.Image
	if v == theme.VariantLight {
		img = canvas.NewImageFromResource(lights[a.name])
	} else {
		img = canvas.NewImageFromResource(darks[a.name])
	}

	img.FillMode = canvas.ImageFillCover
	img.ScaleMode = canvas.ImageScaleFastest
	return img
}

// Name returns a string representation of this design
func (a *Abstract) Name() string {
	return "Abstract " + a.name
}

// All returns a list of all the abstract backgrounds bundled in this package
func All() []*Abstract {
	ret := make([]*Abstract, len(names))
	for i, name := range names {
		ret[i] = &Abstract{name: name}
	}
	return ret
}
