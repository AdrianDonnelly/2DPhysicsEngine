package renderer

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x40, 0x40, 0x40, 0xff})

	for _, circle := range g.placedCircles {
		vector.DrawFilledCircle(screen, circle.X, circle.Y, circle.Radius, circle.Color, true)
	}

}
