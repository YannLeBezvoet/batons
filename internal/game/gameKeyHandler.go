package game

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

func GameKeyHandler(key tcell.Key, carac rune, gameData *GameStruct) int {
	const delay = 200 * time.Millisecond // 100ms
	if key == tcell.KeyEscape {
		return 1
	}
	if key == tcell.KeyLeft {
		if checkDelay(gameData.XCameraTime, delay) {
			gameData.XCameraTime = time.Now()
			gameData.XCamera--
		}
	}
	if key == tcell.KeyRight {
		if checkDelay(gameData.XCameraTime, delay) {
			gameData.XCameraTime = time.Now()
			gameData.XCamera++
		}
	}
	if key == tcell.KeyUp {
		if checkDelay(gameData.YCameraTime, delay) {
			gameData.YCameraTime = time.Now()
			gameData.YCamera--
		}
	}
	if key == tcell.KeyDown {
		if checkDelay(gameData.YCameraTime, delay) {
			gameData.YCameraTime = time.Now()
			gameData.YCamera++
		}
	}
	if carac == 'z' || carac == 'Z' {
		if checkDelay(gameData.YCursorTime, delay) {
			gameData.YCursorTime = time.Now()
			gameData.YCursor--
		}
	}
	if carac == 's' || carac == 'S' {
		if checkDelay(gameData.YCursorTime, delay) {
			gameData.YCursorTime = time.Now()
			gameData.YCursor++
		}
	}
	if carac == 'q' || carac == 'Q' {
		if checkDelay(gameData.XCursorTime, delay) {
			gameData.XCursorTime = time.Now()
			gameData.XCursor--
		}
	}
	if carac == 'd' || carac == 'D' {
		if checkDelay(gameData.XCursorTime, delay) {
			gameData.XCursorTime = time.Now()
			gameData.XCursor++
		}
	}
	return 0
}

func checkDelay(lastTime time.Time, delay time.Duration) bool {
	return time.Since(lastTime) > delay
}
