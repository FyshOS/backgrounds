package abstract

import (
	"embed"

	"fyne.io/fyne/v2"
)

//go:embed "*"
var files embed.FS

var (
	names = []string{"bars", "bow", "cloud", "column", "comb", "cross", "peak", "pine", "rays", "ring", "stripes", "subset"}

	darks  = map[string]fyne.Resource{}
	lights = map[string]fyne.Resource{}
)

func init() {
	for _, name := range names {
		data, _ := files.ReadFile(name + "/svg/dark.svg")
		darks[name] = &fyne.StaticResource{
			StaticName:    name + "_dark.svg",
			StaticContent: data,
		}

		data, _ = files.ReadFile(name + "/svg/lite.svg")
		lights[name] = &fyne.StaticResource{
			StaticName:    name + "_light.svg",
			StaticContent: data,
		}
	}
}
