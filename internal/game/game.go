package game

import (
	"batons/internal/blocks"
	"strconv"

	"github.com/gdamore/tcell/v2"
)

func Game(screen tcell.Screen, gameData GameStruct) {
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

	// Affiche la box de sélection
	DrawSelectionBox(screen)
}

func DrawMap(screen tcell.Screen, gameData GameStruct) {
	for x, y := range gameData.GameMap {
		for yKey, val := range y {
			block, ok := blocks.Get(val)
			if !ok {
				screen.Fini()
				panic("Block not found: " + strconv.Itoa(val))
			}
			style := tcell.StyleDefault.Foreground(block.Color).Background(tcell.ColorBlack)
			screen.SetContent(-gameData.XCamera+x, -gameData.YCamera+yKey, block.Char, nil, style)
		}
	}
}

func DrawStickmen(screen tcell.Screen, gameData GameStruct) {
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	for _, stickman := range gameData.StickManSlice {
		screen.SetContent(-gameData.XCamera+stickman.X, -gameData.YCamera+stickman.Y, '𐀪', nil, style)
	}
}

func DrawSelectionBox(screen tcell.Screen) {
	const selectionBoxSize = 6
	const boxBottomY = 3
	const blockOffsetX = 3
	const blockPosY = 1

	style := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	x, _ := screen.Size()
	blockPosX := x - blockOffsetX
	for i := x - selectionBoxSize; i < x; i++ {
		screen.SetContent(i, boxBottomY, '#', nil, style)
	}
	for i := 0; i < boxBottomY; i++ {
		screen.SetContent(x-selectionBoxSize, i, '#', nil, style)
	}
	screen.SetContent(blockPosX, blockPosY, '█', nil, style)
}
