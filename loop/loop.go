package loop

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"image/color"
	"math"
)

type Particle struct {
	Pt       image.Point
	CenterPt image.Point
	Img      *ebiten.Image
	Color    color.Color
}

func (p *Particle) Draw(screen *ebiten.Image) {
	x := float64(p.Pt.X)
	y := float64(p.Pt.Y)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(p.CenterPt.X), -float64(p.CenterPt.Y))
	op.GeoM.Translate(x, y)
	r, g, b, a := p.Color.RGBA()
	op.ColorM.Scale(r, g, b, a)
	screen.DrawImage(p.Img, op)
}

func NewParticle(x, y int, img *ebiten.Image) *Particle {
	rect := img.Bounds()
	cX := math.Abs(float64(rect.Max.X-rect.Min.X)) / 2
	cY := math.Abs(float64(rect.Max.Y-rect.Min.Y)) / 2
	return &Particle{
		Pt:       image.Pt(x, y),
		CenterPt: image.Pt(int(cX), int(cY)),
		Color:    color.RGBA{0xff, 0x00, 0x00, 0xff},
		Img:      img,
	}
}

type Game struct {
	Width     int
	Height    int
	Particles []*Particle
}

func NewGame(w, h int) *Game {
	img := ebiten.NewImage(50, 50)
	img.Fill(color.White)
	particles := make([]*Particle, 0)
	for i := 0; i < 1; i++ {
		particles = append(particles, NewParticle(w/2, h/2, img))
	}
	return &Game{
		Width:     w,
		Height:    h,
		Particles: particles,
	}
}

func (g *Game) Layout(w, h int) (int, int) {
	return g.Width, g.Height
}

func DebugInfo(screen *ebiten.Image) {
	msg := fmt.Sprintf(`TPS: %0.2f FPS: %0.2f`, ebiten.CurrentTPS(), ebiten.CurrentFPS())
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, particle := range g.Particles {
		particle.Draw(screen)
	}
	DebugInfo(screen)
}

func (g *Game) Update() error {
	return nil
}
