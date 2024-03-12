package tools

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

type DarkTheme struct{}

func (d DarkTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.RGBA{R: 32, G: 32, B: 32, A: 255} // Set the background color
	case theme.ColorNamePrimary:
		return color.RGBA{G: 128, B: 128, A: 255} // Set the primary color
	case theme.ColorNameHyperlink:
		return color.RGBA{R: 192, G: 192, B: 192, A: 255} // Set the text color
	case theme.ColorNameButton:
		return color.RGBA{R: 64, G: 64, B: 64, A: 255} // Set the button color
	case theme.ColorNameDisabledButton:
		return color.RGBA{R: 128, G: 128, B: 128, A: 255} // Set the disabled button color
	default:
		return theme.DefaultTheme().Color(name, variant)
	}
}

func (d DarkTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (d DarkTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (d DarkTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
