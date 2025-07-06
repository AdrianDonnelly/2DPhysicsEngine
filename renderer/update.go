package renderer

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

func (g *Game) Update() error {
	//g.Translate()

	g.mouseX, g.mouseY = ebiten.CursorPosition()

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if !g.isMouseLeftPressed {
			newCircle := PlacedCircle{
				X:      float32(g.mouseX),
				Y:      float32(g.mouseY),
				Color:  color.RGBA{0xff, 0xff, 0x00, 0xff},
				Radius: 5.0,
			}
			g.placedCircles = append(g.placedCircles, newCircle)
			g.isMouseLeftPressed = true
		}
	} else {
		if g.isMouseLeftPressed {
			g.isMouseLeftPressed = false
		}
	}
	return nil
}
