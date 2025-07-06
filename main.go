// Package declaration
package main

// Import statements
import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	dotSize      = 2.0
	moveSpeed    = 1.0
	sideLength   = 100.0
	squareStartX = 50.0
	squareStartY = 50.0
	squareEndX   = squareStartX + sideLength
	squareEndY   = squareStartY + sideLength
)

type MoveState int

const (
	MoveRight MoveState = iota
	MoveDown
	MoveLeft
	MoveUp
)

type Game struct {
	dotX      float32
	dotY      float32
	moveState MoveState
}

func (g *Game) init() {
	g.dotX = 50
	g.dotY = 50
	g.moveState = MoveRight
}

func (g *Game) Update() error {
	switch g.moveState {
	case MoveRight:
		g.dotX += moveSpeed
	case MoveDown:
		g.dotY += moveSpeed
	case MoveLeft:
		g.dotX -= moveSpeed
	case MoveUp:
		g.dotY -= moveSpeed
	}
	switch g.moveState {
	case MoveRight:
		// If dotX has reached or passed the right edge of the square
		if g.dotX >= squareEndX {
			g.dotX = squareEndX // Snap to the exact corner
			g.moveState = MoveDown
		}
	case MoveDown:
		// If dotY has reached or passed the bottom edge
		if g.dotY >= squareEndY {
			g.dotY = squareEndY // Snap to the exact corner
			g.moveState = MoveLeft
		}
	case MoveLeft:
		// If dotX has reached or passed the left edge (moving left, so check if it's <= startX)
		if g.dotX <= squareStartX {
			g.dotX = squareStartX // Snap to the exact corner
			g.moveState = MoveUp
		}
	case MoveUp:
		// If dotY has reached or passed the top edge (moving up, so check if it's <= startY)
		if g.dotY <= squareStartY {
			g.dotY = squareStartY   // Snap to the exact corner
			g.moveState = MoveRight // Loop back to the start
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x40, 0x40, 0x40, 0xff})
	dotColor := color.RGBA{0xff, 0xff, 0x00, 0xff}
	vector.DrawFilledCircle(screen, g.dotX, g.dotY, dotSize, dotColor, false)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

// Func main
func main() {
	game := &Game{}
	game.init() // Initialize the game state
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("2D Physics Engine")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
