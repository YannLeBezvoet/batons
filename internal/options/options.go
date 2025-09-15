package options

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

type OptionsStruct struct {
	MoveCursorLeft  string
	MoveCursorRight string
	MoveCursorUp    string
	MoveCursorDown  string
}

var style = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
var highlightStyle = tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorBlack)

func Options(screen tcell.Screen, optionsAction OptionsAction) {
	mainText := "Options"
	_ = mainText
	menu_buttons := []string{"azerty default", "querty default", "save", "Back"}
	_ = menu_buttons
	width, height := screen.Size()
	// Calculer la position pour centrer le texte
	x := (width - len(mainText)) / 2
	y := height / 2
	// Affiche le texte caractère par caractère
	if optionsAction.Action == None {
		draw(screen, mainText, x, y, false)
		draw(screen, "=========", x-1, y+1, false)
		indicationText := "Use arrow keys to navigate and Enter or Space to select"
		draw(screen, indicationText, (width-len(indicationText))/2, y+2, false)
		for i, button := range menu_buttons {
			isSelected := i == optionsAction.Selected
			draw(screen, button, x, y+3+i, isSelected)
		}
	} else if optionsAction.Action == Save {
		draw(screen, "Settings saved!", x, y, false)
		time.Sleep(500 * time.Millisecond)
	}
}

func draw(screen tcell.Screen, text string, x int, y int, isSelected bool) {
	usedstyle := style
	if isSelected {
		usedstyle = highlightStyle
	}
	for i, r := range text {
		screen.SetContent(x+i, y, r, nil, usedstyle)
	}
}
