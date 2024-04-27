package tools

import (
	"image/color"
	"io/ioutil"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type CustomTheme struct{}

func (d CustomTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	variant = theme.VariantLight
	switch name {
	case theme.ColorNameBackground, theme.ColorNameInputBackground,
		theme.ColorNameOverlayBackground, theme.ColorNameMenuBackground:
		if variant == theme.VariantLight {
			return &color.NRGBA{R: 205, G: 245, B: 253, A: 255}
		}
		return theme.DefaultTheme().Color(name, variant)
	case theme.ColorNameForeground:
		if variant == theme.VariantLight {
			return &color.NRGBA{R: 198, G: 137, B: 198, A: 250}
		}
		return theme.DefaultTheme().Color(name, variant)
	case theme.ColorNamePrimary:
		return &color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xAA}
	case theme.ColorNameButton, theme.ColorNameSelection:
		return &color.NRGBA{R: 137, G: 207, B: 243, A: 0x66}
	case theme.ColorNameFocus:
		return color.RGBA{R: 255, G: 206, B: 243, A: 255}
	default:
		return theme.DefaultTheme().Color(name, variant)
	}
}

func (d CustomTheme) Font(style fyne.TextStyle) fyne.Resource {
	font, err := ioutil.ReadFile("tools/RobotoSlab-Regular.ttf")
	if err != nil {
		log.Fatal(err)
	}

	return &fyne.StaticResource{
		StaticName:    "RobotoSlab-Regular.ttf",
		StaticContent: font,
	}
}

func (d CustomTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (d CustomTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNameText:
		return theme.DefaultTheme().Size(name) + 2
	}

	return theme.DefaultTheme().Size(name)
}
