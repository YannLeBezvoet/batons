package stickman

import (
	"batons/internal/blocks"
	"batons/internal/utils"
	"time"
)

type Stickman struct {
	X                 int
	Y                 int
	Health            int
	XDirection        int //  negatif pour gauche, positif pour droite
	TakeADecisionTime time.Time
	TimeSinceSpawn    time.Time
}

func NewStickman(x, y int) *Stickman {
	return &Stickman{
		X:                 x,
		Y:                 y,
		Health:            100,
		XDirection:        0,
		TakeADecisionTime: time.Now(),
		TimeSinceSpawn:    time.Now(),
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

func (s *Stickman) Update(gameMap map[int]map[int]int) {
	if !s.IsAlive() {
		return
	}
	// Fall of stickman if no ground below
	if gameMap[s.X] == nil {
		s.Move(0, 1) // Move down
	} else if block := blocks.Get(gameMap[s.X][s.Y+1]); !block.IsSolid {
		s.Move(0, 1) // Move down
	}
	// Apply horizontal movement
	s.ApplyMovement(gameMap)
	// Take a decision (move left, right or stay)
	s.TakeADecision(gameMap)
}

func (s *Stickman) TakeADecision(gameMap map[int]map[int]int) {
	// Take a decision every 2 seconds
	block := blocks.Get(gameMap[s.X][s.Y+1])
	if time.Since(s.TakeADecisionTime) > 2*time.Second && block.IsSolid {
		s.TakeADecisionTime = time.Now()
		// Randomly decide to move left, right or stay
		decision := utils.RandRange(0, 3) // 0: left, 1: right, 2: stay
		switch decision {
		case 0:
			s.XDirection = -utils.RandRange(1, 5) // Move left
		case 1:
			s.XDirection = utils.RandRange(1, 5) // Move right
		default:
			s.XDirection = 0 // Stay
		}
	}
}

func (s *Stickman) ApplyMovement(gameMap map[int]map[int]int) {
	blockRight := blocks.Get(gameMap[s.X+1][s.Y])
	blockUpperRight := blocks.Get(gameMap[s.X+1][s.Y-1])
	blockLowerRight := blocks.Get(gameMap[s.X+1][s.Y+1])
	upperBlock := blocks.Get(gameMap[s.X][s.Y-1])
	blockLeft := blocks.Get(gameMap[s.X-1][s.Y])
	blockLowerLeft := blocks.Get(gameMap[s.X-1][s.Y+1])
	lowerBlock := blocks.Get(gameMap[s.X][s.Y+1])
	blockUpperLeft := blocks.Get(gameMap[s.X-1][s.Y-1])

	// Simple horizontal movement
	if s.XDirection > 0 {
		// Move right
		if !blockRight.IsSolid { //if no wall on the right
			s.Move(1, 0)
			s.XDirection--
		} else if !blockLowerRight.IsSolid && !lowerBlock.IsSolid { // if no wall on the down and right (bottom-right diagonal)
			s.Move(1, 1)
			s.XDirection--
		} else if !blockUpperRight.IsSolid && !upperBlock.IsSolid { // if no wall on the up and right (top-right diagonal)
			s.Move(1, -1)
			s.XDirection--
		}
	} else if s.XDirection < 0 {
		// Move left
		if !blockLeft.IsSolid { // if no wall on the left
			s.Move(-1, 0)
			s.XDirection++
		} else if !blockLowerLeft.IsSolid && !lowerBlock.IsSolid { // if no wall on the left and down
			s.Move(-1, 1)
			s.XDirection++
		} else if !blockUpperLeft.IsSolid && !upperBlock.IsSolid { // if no wall on the left and up
			s.Move(-1, -1)
			s.XDirection++
		}
	}
}
