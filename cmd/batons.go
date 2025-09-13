package main

import (
	menu "batons/internal/menu"

	"github.com/gdamore/tcell"
)

func main() {
	// Initialisation de l’écran
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err := screen.Init(); err != nil {
		panic(err)
	}
	defer screen.Fini()
	menu.Menu(screen)
}
