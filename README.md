# Backgrounds

This repository contains a colection of desktop backgrounds
and some simple APIs for including them in Fyne applications.

## Examples

Using the code:

```go
img := backgrounds.Default().Load(
	a.Settings().Theme(), a.Settings().ThemeVariant())
w.SetContent(img)
```

you will fill the window with the default background for the current theme variant.

![](builtin/dark.png)

Or the light variant

![](builtin/light.png)

## Thanks

Many thanks to Joshua Grant for the original design work.