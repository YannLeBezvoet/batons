package game

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

func GameKeyHandler(key tcell.Key, gameData *GameStruct) int {
	const delay = 200 * time.Millisecond // 100ms

	if key == tcell.KeyEscape {
		return 1
	}
	if key == tcell.KeyLeft {
		if checkDelay(gameData.XTime, delay) {
			gameData.XTime = time.Now()
			gameData.X--
		}
	}
	if key == tcell.KeyRight {
		if checkDelay(gameData.XTime, delay) {
			gameData.XTime = time.Now()
			gameData.X++
		}
	}
	if key == tcell.KeyUp {
		if checkDelay(gameData.YTime, delay) {
			gameData.YTime = time.Now()
			gameData.Y--
		}
	}
	if key == tcell.KeyDown {
		if checkDelay(gameData.YTime, delay) {
			gameData.YTime = time.Now()
			gameData.Y++
		}
	}
	return 0
}

func checkDelay(lastTime time.Time, delay time.Duration) bool {
	return time.Since(lastTime) > delay
}
