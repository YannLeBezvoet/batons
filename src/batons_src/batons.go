package main

import (
	"fmt"
	"time"

	"github.com/nsf/termbox-go"
)

func main() {
	// Initialisation de termbox
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		text := fmt.Sprintf("Batons!")
		for i, r := range text {
			termbox.SetCell(i+2, 2, r, termbox.ColorGreen, termbox.ColorDefault)
		}
		termbox.Flush()
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Ch == 'q' || ev.Key == termbox.KeyEsc {
				return
			}
		}
		time.After(50 * time.Millisecond)
	}
}
