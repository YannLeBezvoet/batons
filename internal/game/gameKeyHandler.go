package game

import (
	config "batons/internal/configuration"
	"time"

	"github.com/gdamore/tcell/v2"
)

func GameKeyHandler(key tcell.Key, carac rune, gameData *GameStruct, config config.ConfigStruct) int {
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
	if carac == config.MoveCursorUp || carac == config.MoveCursorUp-32 {
		if checkDelay(gameData.YCursorTime, delay) {
			gameData.YCursorTime = time.Now()
			gameData.YCursor--
		}
	}
	if carac == config.MoveCursorDown || carac == config.MoveCursorDown-32 {
		if checkDelay(gameData.YCursorTime, delay) {
			gameData.YCursorTime = time.Now()
			gameData.YCursor++
		}
	}
	if carac == config.MoveCursorLeft || carac == config.MoveCursorLeft-32 {
		if checkDelay(gameData.XCursorTime, delay) {
			gameData.XCursorTime = time.Now()
			gameData.XCursor--
		}
	}
	if carac == config.MoveCursorRight || carac == config.MoveCursorRight-32 {
		if checkDelay(gameData.XCursorTime, delay) {
			gameData.XCursorTime = time.Now()
			gameData.XCursor++
		}
	}

	if carac == '0' {
		// ajoute un mur à la position du curseur)
		if gameData.GameMap[gameData.XCursor] == nil {
			gameData.GameMap[gameData.XCursor] = make(map[int]int)
		}
		gameData.GameMap[gameData.XCursor][gameData.YCursor] = 1

	}

	return 0
}

func checkDelay(lastTime time.Time, delay time.Duration) bool {
	return time.Since(lastTime) > delay
}
