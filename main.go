package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jairoandre/genuary2022/day01"
	"log"
)

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Spiral")
	game := day01.Init(800, 600, 20)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
