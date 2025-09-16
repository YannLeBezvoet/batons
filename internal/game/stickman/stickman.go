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

func (s *Stickman) TakeDamage(damage int) {
	s.Health -= damage
	if s.Health < 0 {
		s.Health = 0
	}
}

func (s *Stickman) IsAlive() bool {
	return s.Health > 0
}

func (s *Stickman) Heal(amount int) {
	s.Health += amount
	if s.Health > 100 {
		s.Health = 100
	}
}

func (s *Stickman) GetPosition() (int, int) {
	return s.X, s.Y
}

func (s *Stickman) GetHealth() int {
	return s.Health
}

func (s *Stickman) SetPosition(x, y int) {
	s.X = x
	s.Y = y
}

func (s *Stickman) SetHealth(health int) {
	s.Health = health
	if s.Health < 0 {
		s.Health = 0
	} else if s.Health > 100 {
		s.Health = 100
	}
}
