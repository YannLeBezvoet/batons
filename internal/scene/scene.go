package scene

import "github.com/gdamore/tcell"

type Scene struct {
	Width  int
	Height int
	Screen tcell.Screen
}

func NewScene(screen tcell.Screen) *Scene {
	w, h := screen.Size()
	return &Scene{
		Width:  w,
		Height: h,
		Screen: screen,
	}
}

func (s *Scene) Clear() {
	s.Screen.Clear()
}

func (s *Scene) Show() {
	s.Screen.Show()
}

func (s *Scene) SetContent(x, y int, mainc rune, combc []rune, style tcell.Style) {
	s.Screen.SetContent(x, y, mainc, combc, style)
}

func (s *Scene) Size() (int, int) {
	return s.Screen.Size()
}

func (s *Scene) Sync() {
	s.Screen.Sync()
}

func (s *Scene) Fini() {
	s.Screen.Fini()
}
