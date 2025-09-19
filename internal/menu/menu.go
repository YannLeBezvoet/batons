package menu

import (
	"github.com/gdamore/tcell/v2"
)

func Menu(screen tcell.Screen, selected int) {
	text := "Batons"
	secondaryText := "Use arrow keys to navigate and Enter or Space to select"
	menu_buttons := []string{"Start", "Options", "Quit"}
	menu_buttons[selected] = "-> " + menu_buttons[selected]
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
	for i, r := range secondaryText {
		screen.SetContent(i+(width-len(secondaryText))/2, y+2, r, nil, style)
	}
	// Affiche les boutons du menu
	for i, r := range menu_buttons[0] {
		screen.SetContent(i+x, y+3, r, nil, style)
	}
	for i, r := range menu_buttons[1] {
		screen.SetContent(i+x, y+4, r, nil, style)
	}
	for i, r := range menu_buttons[2] {
		screen.SetContent(i+x, y+5, r, nil, style)
	}
}
