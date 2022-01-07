package particles

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"math/rand"
)

type Particle struct {
	x   float64
	y   float64
	img *ebiten.Image
}

func (p *Particle) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.x, p.y)
	screen.DrawImage(p.img, op)
}

func (p *Particle) update() {
	p.x += rand.Float64() - 0.5
	p.y += rand.Float64() - 0.5
}

type Game struct {
	particles []*Particle
	images    []*ebiten.Image
}

func NewColor(r, g, b, a uint8) color.RGBA {
	return color.RGBA{
		R: r,
		G: g,
		B: b,
		A: a,
	}

}

func NewGame() *Game {
	images := make([]*ebiten.Image, 0)
	image := ebiten.NewImage(10, 10)
	image.Fill(NewColor(0xff, 0xff, 0xff, 0xa0))
	images = append(images, image)
	particles := make([]*Particle, 0)
	for i := 0; i < 100; i++ {
		particles = append(particles, &Particle{x: rand.Float64() * 320, y: rand.Float64() * 240, img: images[0]})
	}
	return &Game{
		particles: particles,
	}
}

func (g *Game) Update() error {
	for _, particle := range g.particles {
		particle.update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, particle := range g.particles {
		particle.draw(screen)
	}
}

func (g *Game) Layout(w, h int) (int, int) {
	return 320, 240
}
