package game

import "github.com/gdamore/tcell"

func GameKeyHandler(key tcell.Key) int {
	if key == tcell.KeyEscape {
		return 1
	}
	return 0
}
