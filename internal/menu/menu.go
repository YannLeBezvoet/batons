package menu

import (
	"github.com/gdamore/tcell/v2"
)

func Menu(screen tcell.Screen, selected int) {
	text := "Batons"
	menu_buttons := []string{"Start", "Options", "Quit"}
	// Style simple (blanc sur noir)
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	screen.SetStyle(style)
	// Goroutine pour écouter les événements clavier
	width, height := screen.Size()
	// Calculer la position pour centrer le texte
	x := (width - len(text)) / 2
	y := height / 2
	// Affiche le texte caractère par caractère
	for i, r := range text {
		screen.SetContent(i+x, y, r, nil, style)
	}
	for i := range len(text) + 2 {
		screen.SetContent(i+x-1, y+1, '=', nil, style)
	}
	for i, r := range menu_buttons[0] {
		used_style := style
		if 0 == selected {
			used_style = tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite)
		}
		screen.SetContent(i+x, y+3, r, nil, used_style)
	}
	for i, r := range menu_buttons[1] {
		used_style := style
		if 1 == selected {
			used_style = tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite)
		}
		screen.SetContent(i+x, y+4, r, nil, used_style)
	}
	for i, r := range menu_buttons[2] {
		used_style := style
		if 2 == selected {
			used_style = tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite)
		}
		screen.SetContent(i+x, y+5, r, nil, used_style)
	}
}
