package game

import "github.com/gdamore/tcell"

func GameKeyHandler(key tcell.Key, gameData GameStruct) int {
	if key == tcell.KeyEscape {
		return 1
	}
	return 0
}
