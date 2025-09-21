package blocks

import "github.com/gdamore/tcell/v2"

type Block struct {
	ID      int
	Char    string
	Color   tcell.Color
	IsSolid bool
}
