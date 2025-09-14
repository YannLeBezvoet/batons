package game

import "github.com/gdamore/tcell"

func GameKeyHandler(key tcell.Key, gameData *GameStruct) int {
	if key == tcell.KeyEscape {
		return 1
	}
	if key == tcell.KeyLeft {
		gameData.X--
	}
	if key == tcell.KeyRight {
		gameData.X++
	}
	if key == tcell.KeyUp {
		gameData.Y--
	}
	if key == tcell.KeyDown {
		gameData.Y++
	}
	return 0
}
