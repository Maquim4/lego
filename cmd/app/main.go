package main

import (
	"log"

	"github.com/Maquim4/lego/internal/app"
)

func init() { log.SetFlags(log.Lshortfile | log.LstdFlags) }

func main() {
	legoApp := app.NewApp()
	legoApp.Run()
}
