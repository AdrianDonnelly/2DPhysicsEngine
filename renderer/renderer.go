package renderer

import (
	"image/color"
)

type Game struct {
	dotX                float32
	dotY                float32
	moveState           MoveState
	mouseX              int
	mouseY              int
	isMouseLeftPressed  bool
	hasClickedThisFrame bool
	placedCircles       []PlacedCircle
}
type PlacedCircle struct {
	X      float32
	Y      float32
	Color  color.RGBA
	Radius float32
}

type MoveState int

func (g *Game) Init() {
	g.dotX = 300
	g.dotY = 100
	g.isMouseLeftPressed = false
	g.hasClickedThisFrame = false
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
