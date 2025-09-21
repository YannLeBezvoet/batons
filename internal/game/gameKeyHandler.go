package game

import (
	config "batons/internal/configuration"
	"batons/internal/utils"
	"time"

	"github.com/gdamore/tcell/v2"
)

func GameKeyHandler(key tcell.Key, carac rune, gameData *GameStruct, config config.ConfigStruct) int {
	const delay = 200 * time.Millisecond // 100ms
	if key == tcell.KeyEscape {
		return 1
	}
	if key == tcell.KeyLeft {
		if utils.CheckDelay(gameData.XCameraTime, delay) {
			gameData.XCameraTime = time.Now()
			gameData.XCamera--
		}
	}
	if key == tcell.KeyRight {
		if utils.CheckDelay(gameData.XCameraTime, delay) {
			gameData.XCameraTime = time.Now()
			gameData.XCamera++
		}
	}
	if key == tcell.KeyUp {
		if utils.CheckDelay(gameData.YCameraTime, delay) {
			gameData.YCameraTime = time.Now()
			gameData.YCamera--
		}
	}
	if key == tcell.KeyDown {
		if utils.CheckDelay(gameData.YCameraTime, delay) {
			gameData.YCameraTime = time.Now()
			gameData.YCamera++
		}
	}
	return 0
}
