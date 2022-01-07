package noise

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	time   float64
	scale  int
	width  int
	height int
	halfW  float64
	halfH  float64
	halfS  float64
	Thing  *Thing
}

func Init(w, h, s int) *Game {
	game := Game{time: 0.0, width: w, height: h, scale: s}
	game.halfW = float64(w) / 2.0
	game.halfH = float64(h) / 2.0
	game.halfS = float64(s) / 2.0
	game.Thing = NewThing(0, 0, w, h, 0.02)
	return &game
}

func (g *Game) Update() error {
	g.time += .02
	g.Thing.update(g.time)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Thing.draw(screen)
	msg := fmt.Sprintf(`TPS: %0.2f FPS: %0.2f`, ebiten.CurrentTPS(), ebiten.CurrentFPS())
	ebitenutil.DebugPrint(screen, msg)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.width, g.height
}
