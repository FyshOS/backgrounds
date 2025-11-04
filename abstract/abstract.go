package abstract

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

type Abstract struct {
	name string
}

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

func (a *Abstract) Name() string {
	return "Abstract " + a.name
}

func All() []*Abstract {
	ret := make([]*Abstract, len(names))
	for i, name := range names {
		ret[i] = &Abstract{name: name}
	}
	return ret
}
