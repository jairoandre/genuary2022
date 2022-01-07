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
	Vx  float64
	Vy  float64
	Img *ebiten.Image
}

func (t *Thing) Update(g *Game) {
	if g.time < 25.0 {
		t.X += (rand.Float64() - 0.5) * 2.0
		t.Y += (rand.Float64() - 0.5) * 2.0
	} else {
		dX := t.X - g.halfW
		dY := t.Y - g.halfH
		t.Vx -= dX * 0.0001 * rand.Float64()
		t.Vy -= dY * 0.0001 * rand.Float64()
		t.X += t.Vx
		t.Y += t.Vy
	}
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

func (g *Game) GenerateThings() {
	things := make([]*Thing, 0)
	for y := 0; y < g.height; y += g.scale {
		for x := 0; x < g.width; x += g.scale {
			idx := int(rand.Float64() * float64(len(g.pallet)))
			things = append(things, NewThing(x, y, g.pallet[idx]))
		}
	}
	g.things = things
}

func (g *Game) GenerateOneThing() {
	things := make([]*Thing, 0)
	for j := 0; j < g.height/g.scale; j++ {
		for i := 0; i < g.width/g.scale; i++ {
			things = append(things, NewThing(i*g.scale, j*g.scale, g.pallet[uint(rand.Float64()*float64(len(g.pallet)))]))
		}
	}
	fmt.Println(len(things))
	g.things = things
}

func Init(w, h, s int) *Game {
	game := Game{time: 0.0, width: w, height: h, scale: s}
	game.halfW = float64(w) / 2.0
	game.halfH = float64(h) / 2.0
	game.halfS = float64(s) / 2.0
	pallet := make([]*ebiten.Image, 0)
	for i := 0; i < 10; i++ {
		pallet = append(pallet, NewImg(game.scale, game.scale, RndColor()))
	}
	game.pallet = pallet
	//game.GenerateThings()
	game.GenerateOneThing()
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
		Y:  150,
		Cb: RnUint8() / 3,
		Cr: 255,
	}
	r, g, b, _ := col.RGBA()
	return color.RGBA{
		R: uint8(r), G: uint8(g), B: uint8(b), A: 150,
	}
}
func (t *Thing) draw(screen *ebiten.Image, game *Game, theta float64) {
	ebitenutil.DrawRect(t.Img, t.X, t.Y, 10.0, 2.0, color.White)
}

func (t *Thing) draw1(screen *ebiten.Image, game *Game, theta float64) {
	op := &ebiten.DrawImageOptions{}
	// Translate the thing position (y-axis is negative to invert the coordinates)
	//op.GeoM.Rotate(theta)
	op.GeoM.Translate(t.X, t.Y)
	//op.GeoM.Translate(game.halfW-game.halfS, game.halfH-game.halfS)
	screen.DrawImage(t.Img, op)
}

func (g *Game) Update() error {
	g.time += .2
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, thing := range g.things {
		thing.Update(g)
		thing.draw(screen, g, g.time)
	}
	msg := fmt.Sprintf(`TPS: %0.2f FPS: %0.2f`, ebiten.CurrentTPS(), ebiten.CurrentFPS())
	ebitenutil.DebugPrint(screen, msg)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.width, g.height
}
