package main

import (
	"genuary2022/particles"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Spiral")
	game := particles.NewGame(800.0, 600.0)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
