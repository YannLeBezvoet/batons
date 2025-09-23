package game

import (
	"batons/internal/blocks"
	config "batons/internal/configuration"
	"batons/internal/stickman"
	"batons/internal/utils"
	"time"

	"github.com/gdamore/tcell/v2"
)

func GameMouseHandler(mouseEvent tcell.EventMouse, gameData *GameStruct, config config.ConfigStruct) {
	const delay = 200 * time.Millisecond // 200ms
	x, y := mouseEvent.Position()
	x += gameData.XCamera
	y += gameData.YCamera
	switch mouseEvent.Buttons() {
	case tcell.Button1: // Left click
		if gameData.GameMap[x] == nil {
			gameData.GameMap[x] = make(map[int]int)
		}
		gameData.GameMap[x][y] = gameData.SelectedBlock

	case tcell.Button2: // Right click
		if utils.CheckDelay(gameData.LastRightClickTime, delay) {
			newStickman := stickman.NewStickman(x, y, gameData.GameMap)
			if newStickman != nil {
				gameData.StickManSlice = append(gameData.StickManSlice, newStickman)
			}
			gameData.LastRightClickTime = time.Now()
		}
	case tcell.WheelUp: // Scroll up
		if gameData.SelectedBlock < blocks.RegisterSize-1 {
			gameData.SelectedBlock++
		} else {
			gameData.SelectedBlock = 0
		}
	case tcell.WheelDown: // Scroll down
		if gameData.SelectedBlock > 0 {
			gameData.SelectedBlock--
		} else {
			gameData.SelectedBlock = blocks.RegisterSize - 1
		}
	}
}
