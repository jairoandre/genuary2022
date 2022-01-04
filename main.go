package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jairoandre/genuary2022/day01"
	"log"
)

func main() {
	ebiten.SetWindowSize(500, 500)
	ebiten.SetWindowTitle("Spiral")
	game := day01.Init(500, 500, 50)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
