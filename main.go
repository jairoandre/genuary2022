package main

import (
	"genuary2022/noita"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	scene := noita.NewScene()
	w, h := scene.GetDimensions()
	ebiten.SetWindowSize(w, h)
	ebiten.SetWindowTitle(scene.Title)
	if err := ebiten.RunGame(scene); err != nil {
		log.Fatal(err)
	}
}
