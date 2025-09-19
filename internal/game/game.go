package game

import (
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
)

func Game(screen tcell.Screen, gameData GameStruct) (bool, time.Time) {
	screen.Clear()
	// Style simple (Blanc sur noir)
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	screen.SetStyle(style)
	cursor := '𞢈'
	// Affiche le curseur pendant 500ms au début
	if time.Since(gameData.CursorDrawTime) > 500*time.Millisecond {
		gameData.CursorDrawTime = time.Now()
		gameData.ShowFirstCursor = !gameData.ShowFirstCursor
	}
	if gameData.ShowFirstCursor {
		screen.SetContent(-gameData.XCamera+gameData.XCursor, -gameData.YCamera+gameData.YCursor, cursor, nil, style)
	}
	text := "xCamera: " + strconv.Itoa(gameData.XCamera) + " yCamera: " + strconv.Itoa(gameData.YCamera)
	text += " xCursor: " + strconv.Itoa(gameData.XCursor) + " yCursor: " + strconv.Itoa(gameData.YCursor)
	// Afficher le texte
	for i, r := range text {
		screen.SetContent(i, 0, r, nil, style)
	}

	// Affiche la map
	DrawMap(screen, gameData)
	// Affiche les stickmen
	DrawStickmen(screen, gameData)
	// Affiche le curseur
	if !gameData.ShowFirstCursor {
		screen.SetContent(-gameData.XCamera+gameData.XCursor, -gameData.YCamera+gameData.YCursor, cursor, nil, style)
	}
	// Affiche à l’écran
	screen.Show()
	return gameData.ShowFirstCursor, gameData.CursorDrawTime
}

func DrawMap(screen tcell.Screen, gameData GameStruct) {
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	for x, y := range gameData.GameMap {
		for yKey, val := range y {
			if val == 1 {
				screen.SetContent(-gameData.XCamera+x, -gameData.YCamera+yKey, '█', nil, style)
			}
		}
	}
}

func DrawStickmen(screen tcell.Screen, gameData GameStruct) {
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	for _, stickman := range gameData.StickManSlice {
		screen.SetContent(-gameData.XCamera+stickman.X, -gameData.YCamera+stickman.Y, '𞠻', nil, style)
	}
}
