package renderer

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
	"image/color"
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x40, 0x40, 0x40, 0xff})
	fmt.Println(g.count)
	msg := fmt.Sprintf("Count: %d", g.count)
	text.Draw(screen, msg, basicfont.Face7x13, 10, 20, color.White)

	err := DrawFloor(screen)

	if err != nil {
		return
	}
	for _, circle := range g.placedCircles {
		vector.DrawFilledCircle(screen, circle.X, circle.Y, circle.Radius, circle.Color, true)
	}
	for _, Rect := range g.placedRects {
		vector.DrawFilledRect(screen, Rect.X, Rect.Y, Rect.Width, Rect.Height, color.RGBA{200, 90, 90, 1}, false)
	}

}

func DrawFloor(screen *ebiten.Image) error {
	x := float32(0)
	y := float32(200)
	width := float32(500)
	height := float32(50)
	rectColor := color.RGBA{90, 90, 90, 1}
	vector.DrawFilledRect(screen, x, y, width, height, rectColor, false)
	return nil
}
