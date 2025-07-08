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
	placedCircles       []Circle
	placedRects         []placedRect
	timeStep            float32
	AirDens             float32
	count               int
}
type Circle struct {
	X      float32
	Y      float32
	Mass   float32 //Kg
	Yvel   float32 //vertical vel
	Yaccel float32 //vertical accel
	CDrag  float32 //Coe
	bounce float32
	Area   float32
	theta  float32
	omega  float32
	alpha  float32
	Color  color.RGBA
	Radius float32
}
type placedRect struct {
	X      float32
	Y      float32
	Height float32
	Width  float32
	Mass   float32 //Kg
	Yvel   float32 //vertical vel
	Yaccel float32 //vertical accel
	CDrag  float32 //Coe
	bounce float32
	Area   float32
	theta  float32
	omega  float32
	alpha  float32
	Color  color.RGBA
}
type Vector struct {
	X float32
	Y float32
}

type MoveState int

func (g *Game) Init() {
	g.Height = 200
	g.Width = 400
	g.Gravity = 9.81
	g.timeStep = 0.01
	g.AirDens = 1.2
	g.isMouseLeftPressed = false
	g.hasClickedThisFrame = false
}

func (p *placedRect) Init() {

}
func (v *Vector) Init() {
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
