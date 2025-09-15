package menu

import (
	"github.com/gdamore/tcell/v2"
)

const (
	None = iota
	Start
	Options
	Quit
)

type MenuAction struct {
	Selected int
	Action   int
}

func MenukeyHandler(key tcell.Key, carac rune, selected int, menuSize int) MenuAction {
	if key == tcell.KeyEscape {
		return MenuAction{Selected: selected, Action: Quit}
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
	if key == tcell.KeyEnter || carac == ' ' {
		switch selected {
		case 0:
			// Start
			return MenuAction{Selected: selected, Action: Start}
		case 1:
			// Options
			return MenuAction{Selected: selected, Action: Options}
		case 2:
			// Quit
			return MenuAction{Selected: selected, Action: Quit}
		}
	}
	return MenuAction{Selected: selected, Action: None}
}
