package stickman

type Stickman struct {
	X      int
	Y      int
	Health int
}

func NewStickman(x, y int) *Stickman {
	return &Stickman{
		X:      x,
		Y:      y,
		Health: 100,
	}
}

func (s *Stickman) Move(dx, dy int) {
	s.X += dx
	s.Y += dy
}
