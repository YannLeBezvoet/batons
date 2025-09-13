package menu

import (
	"github.com/gdamore/tcell"
)

func MenukeyHandler(key tcell.Key, selected int, menuSize int) int {
	if key == tcell.KeyEscape {
		quit := make(chan struct{})
		close(quit)
		return 0
	}
	if key == tcell.KeyDown {
		selected = selected + 1
		if selected >= menuSize {
			selected = 0
		}
	}
	if key == tcell.KeyUp {
		selected = selected - 1
		if selected < 0 {
			selected = menuSize - 1
		}
	}
	if key == tcell.KeyEnter || key == ' ' {
		switch selected {
		case 2:
			// Quit
			quit := make(chan struct{})
			close(quit)
		}
	}
	return selected
}
