package game

import "time"

type GameStruct struct {
	XCamera     int
	XCameraTime time.Time
	YCamera     int
	YCameraTime time.Time

	XCursor     int
	YCursor     int
	XCursorTime time.Time
	YCursorTime time.Time
}
