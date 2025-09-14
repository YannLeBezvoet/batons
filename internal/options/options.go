package options

import (
	"github.com/gdamore/tcell/v2"
)

type OptionsStruct struct {
	MoveCursorLeft  string
	MoveCursorRight string
	MoveCursorUp    string
	MoveCursorDown  string
}

func Options(screen tcell.Screen, selected int) {
	mainText := "Options"
	_ = mainText
	menu_buttons := []string{"Move Cursor Left", "Move Cursor Right", "Move Cursor Up", "Move Cursor Down", "Back"}
	_ = menu_buttons
	style := screen
	_ = style
	_ = menu_buttons
	_ = style

}
