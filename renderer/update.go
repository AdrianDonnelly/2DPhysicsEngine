package renderer

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"math"
)

func (g *Game) Update() error {
	//g.Translate()
	for i := range g.placedCircles {
		p := &g.placedCircles[i]
		gravity(g, p)
	}

	g.mouseX, g.mouseY = ebiten.CursorPosition()
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if !g.isMouseLeftPressed {
			Radius := 5.0
			newCircle := PlacedCircle{
				X:      float32(g.mouseX),
				Y:      float32(g.mouseY),
				Mass:   0.1,
				Yvel:   0.00,
				CDrag:  0.47,
				bounce: -0.5,
				Yaccel: 0,
				Color:  color.RGBA{0xff, 0xff, 0x00, 0xff},
				Radius: float32(Radius),
				Area:   float32(math.Pi * Radius * Radius / 10000),
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
