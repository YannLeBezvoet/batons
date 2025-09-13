package game

import (
	"strconv"
	"time"

	"github.com/gdamore/tcell"
)

func Game(screen tcell.Screen) {
	x := 0
	y := 0
	text := "x: " + strconv.Itoa(x) + " y: " + strconv.Itoa(y)
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
				if ev.Key() == tcell.KeyEscape {
					return
				}
			case *tcell.EventResize:
				screen.Sync()
			}
		}
	}()
	// Boucle principale pour afficher le message
	for {
		select {
		case <-quit:
			return
		default:
			screen.Clear()
			// Afficher le texte
			for i, r := range text {
				screen.SetContent(i, 0, r, nil, style)
			}
			// Affiche à l’écran
			screen.Show()
			time.Sleep(50 * time.Millisecond)
		}
	}
}
