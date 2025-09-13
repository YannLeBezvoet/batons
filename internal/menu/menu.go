package menu

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell"
)

func Menu(screen tcell.Screen) {

	quit := make(chan struct{})

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
			// Efface l’écran
			screen.Clear()

			// Prépare le texte
			text := fmt.Sprintf("Batons, appuyer sur 'q' ou 'ESC' pour quitter. Heure: %s", time.Now().Format("15:04:05"))

			// Affiche le texte caractère par caractère
			style := tcell.StyleDefault.Foreground(tcell.ColorGreen)
			for i, r := range text {
				screen.SetContent(i+2, 2, r, nil, style)
			}

			// Affiche à l’écran
			screen.Show()

			// Attente 50 milisecondes
			time.Sleep(50 * time.Millisecond)
		}
	}
}
