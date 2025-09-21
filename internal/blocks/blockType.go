package blocks

import "github.com/gdamore/tcell/v2"

type Block struct {
	ID      int
	Char    rune
	Color   tcell.Color
	IsSolid bool
}
