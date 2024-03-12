package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/Maquim4/lego/internal/renderer"
)

func main() {
	a := app.NewWithID("go.lego.quiz")
	w := a.NewWindow("<Enjoy doing tests3>")
	w.SetMaster()
	w.CenterOnScreen()

	rend := renderer.NewMainRenderer()
	rend.Render(w)

	w.Resize(fyne.NewSize(640, 460))
	w.ShowAndRun()
}
