package backgrounds

import (
	"github.com/FyshOS/backgrounds/builtin"

	"fyne.io/fyne/v2"
)

type Background interface {
	Name() string
	Load(t fyne.Theme, v fyne.ThemeVariant) fyne.CanvasObject
}

func Default() Background {
	return builtin.Default()
}
