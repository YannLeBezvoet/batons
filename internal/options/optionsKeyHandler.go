package options

import (
	configuration "batons/internal/configuration"
	"time"

	"github.com/gdamore/tcell/v2"
)

type OptionsAction struct {
	Selected int
	Action   int
}

const (
	None = iota
	azertyDefault
	quertyDefault
	MoveCursorLeft
	MoveCursorRight
	MoveCursorUp
	MoveCursorDown
	Save
	Quit
)

func OptionsKeyHandler(key tcell.Key, carac rune, optionsData *OptionsStruct, selected int, menuSize int, config *configuration.ConfigStruct) OptionsAction {
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
	if key == tcell.KeyEnter || carac == ' ' {
		switch selected {
		case 0:
			config.MoveCursorLeft = 'q'
			config.MoveCursorRight = 'd'
			config.MoveCursorUp = 'z'
			config.MoveCursorDown = 's'
		case 1:
			config.MoveCursorLeft = 'a'
			config.MoveCursorRight = 'd'
			config.MoveCursorUp = 'w'
			config.MoveCursorDown = 's'
		case 2:
			// Save settings to config
			if optionsData != nil {
				configuration.SaveConfig("config.json", *config)
				return OptionsAction{Selected: selected, Action: Save}

			}
		case 3:
			return OptionsAction{Selected: 0, Action: Quit}
		}

	}
	return OptionsAction{Selected: selected, Action: None}
}
