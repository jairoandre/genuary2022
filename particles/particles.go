package particles

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"math/rand"
)

type Particle struct {
	x float64
	y float64
}

func (p *Particle) draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, p.x, p.y, 1.0, 10.0, color.White)
}

type Game struct {
	particles []*Particle
}

func NewGame() *Game {
	particles := make([]*Particle, 0)
	for i := 0; i < 100; i++ {
		particles = append(particles, &Particle{x: rand.Float64() * 320, y: rand.Float64() * 240})
	}
	return &Game{
		particles: particles,
	}
}

func (g *Game) Update() error {
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
