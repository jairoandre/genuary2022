package main

import (
	"genuary2022/particles"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	ebiten.SetWindowSize(300, 240)
	ebiten.SetWindowTitle("Spiral")
	game := particles.NewGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
