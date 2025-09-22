package blocks

import "github.com/gdamore/tcell/v2"

var registry = map[int]Block{}
var RegisterSize int

// Register ajoute un bloc au registre
func Register(b Block) {
	registry[b.ID] = b
}

// Get retourne un bloc par son ID
func Get(id int) Block {
	b, ok := registry[id]
	if !ok {
		return registry[Air] // air
	}
	return b
}

func Init() {
	Register(Block{ID: Air, Char: ' ', Color: tcell.ColorBlack, IsSolid: false})  // air
	Register(Block{ID: Stone, Char: '█', Color: tcell.ColorWhite, IsSolid: true}) // stone
	Register(Block{ID: Dirt, Char: '▒', Color: tcell.ColorBrown, IsSolid: true})  // dirt
	Register(Block{ID: Grass, Char: '▓', Color: tcell.ColorGreen, IsSolid: true}) // grass
	RegisterSize = len(registry)
}
