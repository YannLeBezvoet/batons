package options

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

func OptionsKeyHandler(key tcell.Key, carac rune, optionsData *OptionsStruct) int {
	const delay = 200 * time.Millisecond // 100ms
	if key == tcell.KeyEscape {
		return 1
	}
	return 0
}
