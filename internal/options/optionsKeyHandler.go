package options

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

type OptionsAction struct {
	Selected int
	Action   int
}

const (
	None = iota
	Quit
)

func OptionsKeyHandler(key tcell.Key, carac rune, optionsData *OptionsStruct, selected int, menuSize int) OptionsAction {
	const delay = 200 * time.Millisecond // 100ms
	if key == tcell.KeyEscape {
		return OptionsAction{Selected: 0, Action: Quit}
	}
	if key == tcell.KeyDown {
		selected++
		if selected >= menuSize {
			selected = 0
		}
	}
	if key == tcell.KeyUp {
		selected--
		if selected < 0 {
			selected = menuSize - 1
		}
	}
	return OptionsAction{Selected: selected, Action: 0}
}
