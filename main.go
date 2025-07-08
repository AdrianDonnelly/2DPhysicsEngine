package main

import (
	"2DPhysicsEngine/renderer"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	game := &renderer.Game{}
	game.Init()
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("2D Physics Engine")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
