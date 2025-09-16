package game

import (
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
)

func Game(screen tcell.Screen, gameData GameStruct) {
	text := "xCamera: " + strconv.Itoa(gameData.XCamera) + " yCamera: " + strconv.Itoa(gameData.YCamera)
	text += " xCursor: " + strconv.Itoa(gameData.XCursor) + " yCursor: " + strconv.Itoa(gameData.YCursor)
	// Style simple (blanc sur noir)
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	screen.SetStyle(style)
	// Boucle principale pour afficher le message

	screen.Clear()
	// Afficher le texte
	for i, r := range text {
		screen.SetContent(i, 0, r, nil, style)
	}

	// Affiche la map
	DrawMap(screen, gameData)
	// Affiche le curseur
	cursor := '*'
	if gameData.GameMap[gameData.XCursor][gameData.YCursor] == 1 {
		cursor = '▓'
	}
	screen.SetContent(-gameData.XCamera+gameData.XCursor, -gameData.YCamera+gameData.YCursor, cursor, nil, style)
	// Affiche à l’écran
	screen.Show()
	time.Sleep(50 * time.Millisecond)
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
