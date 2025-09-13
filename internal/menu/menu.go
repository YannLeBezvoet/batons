package menu

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell"
)

func Menu(screen tcell.Screen) {
	text := "Batons"
	menu_buttons := []string{"Start", "Options", "Quit"}
	selected := 0
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
				if ev.Key() == tcell.KeyDown {
					selected = selected + 1
					if selected >= len(menu_buttons) {
						selected = 0
					}
				}
				if ev.Key() == tcell.KeyUp {
					selected = selected - 1
					if selected < 0 {
						selected = len(menu_buttons) - 1
					}
				}
				if ev.Key() == tcell.KeyEnter || ev.Rune() == ' ' {
					switch selected {
					case 0:
						// Start
						// Juste un message pour l’instant
						text = "Game Started! Press 'q' to quit."
					case 1:
						// Options
						text = "Options Selected! Press 'q' to quit."
					case 2:
						// Quit
						close(quit)
						return
					}
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
			// Affiche à l’écran
			screen.Show()

			// Attente 50 milisecondes
			time.Sleep(50 * time.Millisecond)
		}
	}
}
