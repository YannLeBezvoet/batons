package menu

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell"
)

func Menu(screen tcell.Screen) {
	text := "Batons"
	menu_buttons := []string{"Start", "Options", "Quit"}
	// Obtenir la taille de l’écran
	quit := make(chan struct{})
	// Style simple (blanc sur noir)
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	screen.SetStyle(style)
	// Goroutine pour écouter les événements clavier
	go func() {
		for {
			ev := screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape || ev.Rune() == 'q' {
					close(quit)
					return
				}
			case *tcell.EventResize:
				screen.Sync()
			}
		}
	}()

	for {
		select {
		case <-quit:
			return
		default:
			width, height := screen.Size()
			// Calculer la position pour centrer le texte
			x := (width - len(text)) / 2
			y := height / 2
			// Efface l’écran
			screen.Clear()

			// Prépare le texte
			text := fmt.Sprintf("%s", text)

			// Affiche le texte caractère par caractère
			for i, r := range text {
				screen.SetContent(i+x, y, r, nil, style)
			}
			for i := range len(text) + 2 {
				screen.SetContent(i+x-1, y+1, '=', nil, style)
			}
			for i, r := range menu_buttons[0] {
				screen.SetContent(i+x, y+3, r, nil, style)
			}
			for i, r := range menu_buttons[1] {
				screen.SetContent(i+x, y+4, r, nil, style)
			}
			for i, r := range menu_buttons[2] {
				screen.SetContent(i+x, y+5, r, nil, style)
			}
			// Affiche à l’écran
			screen.Show()

			// Attente 50 milisecondes
			time.Sleep(50 * time.Millisecond)
		}
	}
}
