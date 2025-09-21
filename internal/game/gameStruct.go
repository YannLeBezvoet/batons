package game

import (
	"batons/internal/stickman"
	"time"
)

type GameStruct struct {
	XCamera     int
	XCameraTime time.Time
	YCamera     int
	YCameraTime time.Time

	CursorDrawTime  time.Time
	ShowFirstCursor bool

	GameMap            map[int]map[int]int
	EnterTime          time.Time
	StickManSlice      []*stickman.Stickman
	LastRightClickTime time.Time
}
