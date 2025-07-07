package renderer

import (
	"image/color"
)

type Game struct {
	Height              float32
	Width               float32
	Gravity             float32
	moveState           MoveState
	mouseX              int
	mouseY              int
	isMouseLeftPressed  bool
	hasClickedThisFrame bool
	placedCircles       []PlacedCircle
	timeStep            float32
	rho                 float32
}
type PlacedCircle struct {
	X      float32
	Y      float32
	V      float32
	M      float32
	vy     float32 //vertical vel
	ay     float32 //vertical accel
	C_d    float32
	e      float32
	A      float32
	Color  color.RGBA
	Radius float32
}

type MoveState int

func (g *Game) Init() {
	g.Height = 400
	g.Width = 400
	g.Gravity = 9.81
	g.timeStep = 1.0 / 60.0
	g.rho = 1.2
	g.isMouseLeftPressed = false
	g.hasClickedThisFrame = false
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
