package main

import (
	"log"
	"time"

	"batons/internal/game"
	"batons/internal/menu"

	"github.com/gdamore/tcell"
)

type AppState int

const (
	StateMenu AppState = iota
	StateGame
	StateOptions
)

func main() {
	selected := 0
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	defer screen.Fini()

	state := StateMenu

	for {
		start := time.Now()
		screen.Clear()

		switch state {
		case StateMenu:
			menu.Menu(screen, selected)
		case StateGame:
			game.Game(screen)
		}

		screen.Show()

		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			if state == StateMenu {
				selected = menu.MenukeyHandler(ev.Key(), selected, 3)
			}
		case *tcell.EventResize:
			screen.Sync()
		}

		elapsed := time.Since(start)
		if elapsed < 50*time.Millisecond {
			time.Sleep(50*time.Millisecond - elapsed)
		}
	}
}
