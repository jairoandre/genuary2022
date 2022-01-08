package particles

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ojrac/opensimplex-go"
	"image"
	"image/color"
	"math/rand"
)

const (
	particleSize     = 2
	halfParticleSize = particleSize / 2
)

type Particle struct {
	x      float64
	y      float64
	rot    float64
	rotVel float64
	img    *ebiten.Image
}

func (p *Particle) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfParticleSize, -halfParticleSize)
	op.GeoM.Rotate(p.rot)
	op.GeoM.Translate(p.x, p.y)
	screen.DrawImage(p.img, op)
}

func (p *Particle) update(game *Game) {
	sn := 0.03
	p.rot += p.rotVel
	p.x += 2 * (game.noise.Eval3(sn*p.x, sn*p.y, 0.0) - 0.5)
	p.y += 2 * (game.noise.Eval3(sn*p.x, sn*p.y, 1.0) - 0.5)
}

func RndOne() float64 {
	return 2 * (rand.Float64() - 0.5)
}

type Game struct {
	time      float64
	width     float64
	height    float64
	particles []*Particle
	noise     opensimplex.Noise
	images    []*ebiten.Image
	prev      *ebiten.Image
}

func NewColor(r, g, b, a uint8) color.RGBA {
	return color.RGBA{
		R: r,
		G: g,
		B: b,
		A: a,
	}

}

func NewGame(w, h float64) *Game {
	images := make([]*ebiten.Image, 0)
	image := ebiten.NewImage(particleSize, particleSize)
	image.Fill(NewColor(0xff, 0xff, 0xff, 0x0a))
	images = append(images, image)
	particles := make([]*Particle, 0)
	noise := opensimplex.NewNormalized(934)
	for i := 0; i < 10000; i++ {
		particles = append(particles, &Particle{
			x:      RndOne() * w,
			y:      RndOne() * h,
			rotVel: RndOne() * 0.4,
			img:    images[0]})
	}
	return &Game{
		width:     w,
		height:    h,
		noise:     noise,
		particles: particles,
	}
}

func (g *Game) Update() error {
	for _, particle := range g.particles {
		particle.update(g)
	}
	g.time += 0.5
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.prev != nil {
		op := &ebiten.DrawImageOptions{}
		screen.DrawImage(g.prev, op)
	}
	for _, particle := range g.particles {
		particle.draw(screen)
	}
	prev := screen.SubImage(image.Rect(0, 0, int(g.width), int(g.height)))
	g.prev = ebiten.NewImageFromImage(prev)
}

func (g *Game) Layout(w, h int) (int, int) {
	return int(g.width), int(g.height)
}
