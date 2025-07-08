package renderer

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

func (g *Game) Update() error {
	for i := range g.placedRects {
		p := &g.placedRects[i]
		gravity(g, p)
	}

	g.mouseX, g.mouseY = ebiten.CursorPosition()

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if !g.isMouseLeftPressed {
			newRect := placedRect{
				X:      float32(g.mouseX),
				Y:      float32(g.mouseY),
				Height: 10,
				Width:  15,
				Mass:   0.1,
				Yvel:   0.01,
				CDrag:  0.47,
				bounce: -0.5,
				Yaccel: 0.0,
				Color:  color.RGBA{0xff, 0xff, 0x00, 0xff},
			}
			g.count++
			g.placedRects = append(g.placedRects, newRect)
			g.isMouseLeftPressed = true
		}
	} else {
		g.isMouseLeftPressed = false
	}

	return nil
}
