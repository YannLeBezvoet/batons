package game

import (
	"batons/internal/blocks"
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
)

func Game(screen tcell.Screen, gameData GameStruct) (bool, time.Time) {
	// Style simple (Blanc sur noir)
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	screen.SetStyle(style)

	text := "xCamera: " + strconv.Itoa(gameData.XCamera) + " yCamera: " + strconv.Itoa(gameData.YCamera)
	// Afficher le texte
	for i, r := range text {
		screen.SetContent(i, 0, r, nil, style)
	}

	// Affiche la map
	DrawMap(screen, gameData)
	// Update stickmen
	for _, stickman := range gameData.StickManSlice {
		stickman.Update(gameData.GameMap)
	}
	// Affiche les stickmen
	DrawStickmen(screen, gameData)

	// Affiche à l’écran
	return gameData.ShowFirstCursor, gameData.CursorDrawTime
}

func DrawMap(screen tcell.Screen, gameData GameStruct) {
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	for x, y := range gameData.GameMap {
		for yKey, val := range y {
			if val == blocks.Stone {
				screen.SetContent(-gameData.XCamera+x, -gameData.YCamera+yKey, '█', nil, style)
			}
		}
	}
}

func DrawStickmen(screen tcell.Screen, gameData GameStruct) {
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	for _, stickman := range gameData.StickManSlice {
		screen.SetContent(-gameData.XCamera+stickman.X, -gameData.YCamera+stickman.Y, '𐀪', nil, style)
	}
}
