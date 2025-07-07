package renderer

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x40, 0x40, 0x40, 0xff})

	err := DrawFloor(screen)

	if err != nil {
		return
	}
	for _, circle := range g.placedCircles {
		vector.DrawFilledCircle(screen, circle.X, circle.Y, circle.Radius, circle.Color, true)
	}

}

func DrawFloor(screen *ebiten.Image) error {
	x := float32(0)
	y := float32(200)
	width := float32(500)
	height := float32(50)
	rectColor := color.RGBA{0xFF, 0x00, 0x00, 0xFF}
	vector.DrawFilledRect(screen, x, y, width, height, rectColor, false)
	return nil
}
