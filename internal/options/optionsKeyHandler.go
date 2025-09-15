package options

import (
	configuration "batons/internal/configuration"
	"time"

	"github.com/gdamore/tcell/v2"
)

type OptionsAction struct {
	Selected int
	Action   int
	Waiting  bool
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

func OptionsKeyHandler(key tcell.Key, carac rune, optionsAction *OptionsAction, menuSize int, config configuration.ConfigStruct) (OptionsAction, configuration.ConfigStruct) {
	const delay = 200 * time.Millisecond // 100ms
	optionsAction.Waiting = false
	if optionsAction.Action == Save {
		optionsAction.Action = Quit
		return *optionsAction, config
	}
	if key == tcell.KeyEscape {
		return OptionsAction{Selected: 0, Action: Quit}, config
	}
	if key == tcell.KeyDown {
		optionsAction.Selected++
		if optionsAction.Selected >= menuSize {
			optionsAction.Selected = 0
		}
	}
	if key == tcell.KeyUp {
		optionsAction.Selected--
		if optionsAction.Selected < 0 {
			optionsAction.Selected = menuSize - 1
		}
	}
	if key == tcell.KeyEnter || carac == ' ' {
		switch optionsAction.Selected {
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
			configuration.SaveConfig("config.json", config)
			return OptionsAction{Selected: optionsAction.Selected, Action: Save, Waiting: true}, config

		case 3:
			return OptionsAction{Selected: 0, Action: Quit, Waiting: false}, config
		}

	}
	return OptionsAction{Selected: optionsAction.Selected, Action: None, Waiting: false}, config
}
