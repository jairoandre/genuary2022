package day01

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"math"
	"math/rand"
)

var (
	img  *ebiten.Image
	game Game
)

type Thing struct {
	X float64
	Y float64
}

type Game struct {
	time   float64
	scale  int
	width  int
	height int
	halfW  float64
	halfH  float64
	halfS  float64
}

func Init(w, h, s int) *Game {
	game = Game{time: 0.0, width: w, height: h, scale: s}
	game.halfW = float64(w) / 2.0
	game.halfH = float64(h) / 2.0
	game.halfS = float64(s) / 2.0
	img = ebiten.NewImage(s, s)
	col := color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	img.Fill(col)
	return &game
}

func NewImg(w, h int, r, g, b, a uint8) *ebiten.Image {
	img := ebiten.NewImage(w, h)
	img.Fill(color.RGBA{R: r, G: g, B: b, A: a})
	return img
}

func RnUint8() uint8 {
	return uint8(rand.Float64() * 255)
}

func RndColor() color.RGBA {
	return color.RGBA{
		R: RnUint8(),
		G: RnUint8(),
		B: RnUint8(),
		A: RnUint8(),
	}
}

func (t *Thing) draw(screen *ebiten.Image, game *Game, theta float64) {
	op := &ebiten.DrawImageOptions{}
	// Translate the thing position (y-axis is negative to invert the coordinates)
	op.GeoM.Translate(t.X, -t.Y)
	op.GeoM.Rotate(theta)

	op.GeoM.Translate(game.halfW-game.halfS, game.halfH-game.halfS)
	screen.DrawImage(img, op)
}

func (g *Game) Update() error {
	g.time += .02
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	theta := math.Sin(g.time*60.0*0.001) * math.Pi * 2.0 * 4.0
	for y := 0; y < g.height; y++ {
		y := float64(y)
		for i := 0; i < 10; i++ {
			i := float64(i)
			rot := i + 1.0 + y*10.0
			thing := Thing{math.Sin(y*0.1+theta+i*2.0) * 100., y}
			thing.draw(screen, g, theta*0.001*rot)
		}
	}
	msg := fmt.Sprintf(`TPS: %0.2f FPS: %0.2f`, ebiten.CurrentTPS(), ebiten.CurrentFPS())
	ebitenutil.DebugPrint(screen, msg)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.width, g.height
}
