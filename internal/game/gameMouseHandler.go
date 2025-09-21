package game

import (
	config "batons/internal/configuration"
	"batons/internal/stickman"
	"batons/internal/utils"
	"time"

	"github.com/gdamore/tcell/v2"
)

func GameMouseHandler(mouseEvent tcell.EventMouse, gameData *GameStruct, config config.ConfigStruct) {
	const delay = 200 * time.Millisecond // 200ms
	switch mouseEvent.Buttons() {
	case tcell.Button1: // Left click
		x, y := mouseEvent.Position()
		if gameData.GameMap[x] == nil {
			gameData.GameMap[x] = make(map[int]int)
		}
		gameData.GameMap[x][y] = 1

	case tcell.Button2: // Right click
		if utils.CheckDelay(gameData.LastRightClickTime, delay) {
			x, y := mouseEvent.Position()
			newStickman := stickman.NewStickman(x, y)
			gameData.StickManSlice = append(gameData.StickManSlice, newStickman)
			gameData.LastRightClickTime = time.Now()
		}
	}
}
