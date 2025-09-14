package main

import (
	"log"
	"time"

	"batons/internal/game"
	"batons/internal/menu"

	"github.com/gdamore/tcell/v2"
)

type AppState int

const (
	StateMenu AppState = iota
	StateGame
	StateOptions
)

func main() {
	menuAction := menu.MenuAction{Selected: 0, Action: menu.None}
	gameData := game.GameStruct{XCamera: 0, YCamera: 0, XCursor: 1, YCursor: 1}
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
			menu.Menu(screen, menuAction.Selected)
		case StateGame:
			game.Game(screen, gameData)
		}
		go eventListener(screen, &state, &menuAction, &gameData)

		if menuAction.Action == menu.Quit {
			quit := make(chan struct{})
			close(quit)
			return
		}
		screen.Show()

		elapsed := time.Since(start)
		if elapsed < 50*time.Millisecond {
			time.Sleep(50*time.Millisecond - elapsed)
		}
	}
}

func eventListener(screen tcell.Screen, state *AppState, menuAction *menu.MenuAction, gameData *game.GameStruct) {
	ev := screen.PollEvent()
	switch ev := ev.(type) {
	case *tcell.EventKey:
		if *state == StateMenu {
			*menuAction = menu.MenukeyHandler(ev.Key(), menuAction.Selected, 3)
			if menuAction.Action == menu.Start {
				*state = StateGame
				*menuAction = menu.MenuAction{Selected: 0, Action: menu.None}
			}
		}
		if *state == StateGame {
			gameAction := game.GameKeyHandler(ev.Key(), ev.Rune(), gameData)
			if gameAction == 1 {
				*state = StateMenu
				*menuAction = menu.MenuAction{Selected: 0, Action: menu.None}
				gameAction = 0
			}
		}
	case *tcell.EventResize:
		screen.Sync()
	}
}
