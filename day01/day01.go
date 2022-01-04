package day01

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"math/rand"
)

type Thing struct {
	X   float64
	Y   float64
	Img *ebiten.Image
}

func NewThing(x, y int, img *ebiten.Image) *Thing {
	return &Thing{
		X:   float64(x),
		Y:   float64(y),
		Img: img,
	}
}

type Game struct {
	time   float64
	scale  int
	width  int
	height int
	halfW  float64
	halfH  float64
	halfS  float64
	things []*Thing
	pallet []*ebiten.Image
}

func Init(w, h, s int) *Game {
	game := Game{time: 0.0, width: w, height: h, scale: s}
	game.halfW = float64(w) / 2.0
	game.halfH = float64(h) / 2.0
	game.halfS = float64(s) / 2.0
	pallet := make([]*ebiten.Image, 0)
	for i := 0; i < 100; i++ {
		pallet = append(pallet, NewImg(game.scale, game.scale, RndColor()))
	}
	game.pallet = pallet
	things := make([]*Thing, 0)
	for y := 0; y < game.height; y += game.scale {
		for x := 0; x < game.width; x += game.scale {
			idx := int(rand.Float64() * float64(len(pallet)))
			things = append(things, NewThing(x, y, pallet[idx]))
		}
	}
	game.things = things
	fmt.Println(len(things))
	return &game
}

func NewImg(w, h int, col color.RGBA) *ebiten.Image {
	img := ebiten.NewImage(w, h)
	img.Fill(col)
	return img
}

func RnUint8() uint8 {
	return uint8(rand.Float64() * 255)
}

func RndColor() color.RGBA {
	col := color.YCbCr{
		Y:  RnUint8(),
		Cb: RnUint8(),
		Cr: RnUint8(),
	}
	r, g, b, a := col.RGBA()
	return color.RGBA{
		R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a),
	}
}

func (t *Thing) draw(screen *ebiten.Image, game *Game, theta float64) {
	op := &ebiten.DrawImageOptions{}
	// Translate the thing position (y-axis is negative to invert the coordinates)
	op.GeoM.Translate(t.X, t.Y)
	op.GeoM.Rotate(theta)
	//op.GeoM.Translate(game.halfW-game.halfS, game.halfH-game.halfS)
	screen.DrawImage(t.Img, op)
}

func (g *Game) Update() error {
	g.time += .02
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, thing := range g.things {
		thing.draw(screen, g, 0.01*g.time)
	}
	msg := fmt.Sprintf(`TPS: %0.2f FPS: %0.2f`, ebiten.CurrentTPS(), ebiten.CurrentFPS())
	ebitenutil.DebugPrint(screen, msg)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.width, g.height
}
