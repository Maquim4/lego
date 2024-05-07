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
	return theme.DefaultTheme().Color(name, variant)
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
