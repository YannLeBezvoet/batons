package main

import (
	"fmt"

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
	}
}
