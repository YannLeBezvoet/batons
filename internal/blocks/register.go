package blocks

var registry = map[int]Block{}

// Register ajoute un bloc au registre
func Register(b Block) {
	registry[b.ID] = b
}

// Get retourne un bloc par son ID
func Get(id int) (Block, bool) {
	b, ok := registry[id]
	return b, ok
}
