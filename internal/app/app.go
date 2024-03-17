package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/Maquim4/lego/internal/app/pages"
)

const ID = "github.com/Maquim4/lego"

type App struct {
	fapp fyne.App
	fwin fyne.Window
}

func NewApp() *App {
	a := app.NewWithID(ID)

	w := a.NewWindow("<Enjoy,live, go test3>")
	w.SetMaster()
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(640, 460))

	return &App{fapp: a, fwin: w}
}

func (a *App) Run() {
	home := pages.NewMenu()
	home.Render(a.fwin)
	a.fwin.ShowAndRun()
}
