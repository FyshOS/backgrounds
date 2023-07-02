package builtin

import (
	_ "embed"

	"fyne.io/fyne/v2"
)

//go:embed dark.png
var backgroundDark []byte

//go:embed light.png
var backgroundLight []byte

var resourceDarkPng = &fyne.StaticResource{
	StaticName:    "dark.png",
	StaticContent: backgroundDark,
}

var resourceLightPng = &fyne.StaticResource{
	StaticName:    "light.png",
	StaticContent: backgroundLight,
}
