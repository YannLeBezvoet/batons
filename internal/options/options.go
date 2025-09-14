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

var style = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
var highlightStyle = tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorBlack)

func Options(screen tcell.Screen, selected int) {
	mainText := "Options"
	_ = mainText
	menu_buttons := []string{"Move Cursor Left", "Move Cursor Right", "Move Cursor Up", "Move Cursor Down", "Back"}
	_ = menu_buttons
	width, height := screen.Size()
	// Calculer la position pour centrer le texte
	x := (width - len(mainText)) / 2
	y := height / 2
	// Affiche le texte caractère par caractère
	draw(screen, mainText, x, y, false)
	draw(screen, "=========", x-1, y+1, false)
	indicationText := "Use arrow keys to navigate and Enter to select"
	draw(screen, indicationText, (width-len(indicationText))/2, y+2, false)
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
