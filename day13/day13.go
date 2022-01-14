package day13

import (
	"fmt"
	"genuary2022/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mazznoer/colorgrad"
	"image/color"
	"math"
	"math/rand"
)

const (
	lifeSpan = 10
	scale    = 2
	width    = 800
	height   = 80
	wScaled  = width / scale
	hScaled  = height / scale
)

type Scene struct {
	Width        float64
	Height       float64
	Title        string
	Img          *ebiten.Image
	Particles    []*Particle
	FireGradient colorgrad.Gradient
	GifWriter    *utils.GifWriter
}

type Particle struct {
	Pos      utils.Point
	Gradient colorgrad.Gradient
	Img      *ebiten.Image
	LifeSpan float64
}

func NewParticle(x, y float64, gradient colorgrad.Gradient, img *ebiten.Image) *Particle {
	return &Particle{
		Pos:      utils.Pt(x, y),
		Gradient: gradient,
		Img:      img,
		LifeSpan: lifeSpan,
	}
}

func (p *Particle) Update() {
	rng := int(math.Round(rand.Float64())*3.0) & 3
	p.LifeSpan -= float64(rng & 1)
	p.Pos.Y -= 1
	if rand.Float64() < 0.4 {
		p.Pos.X -= 1
	}
}

func (p *Particle) Draw(screen *ebiten.Image) {
	ratio := p.LifeSpan / lifeSpan
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(p.Pos.X*scale, p.Pos.Y*scale)
	r, g, b, a := utils.NormalizeColor(p.Gradient.At(ratio))
	op.ColorM.Scale(r, g, b, a)
	screen.DrawImage(p.Img, op)
}

func NewScene() *Scene {
	img := ebiten.NewImage(1, 1)
	img.Fill(color.White)
	gradient := colorgrad.Inferno()
	particles := make([]*Particle, 0)
	for x := 0; x < wScaled; x++ {
		particles = append(particles, NewParticle(float64(x), hScaled-1, gradient, img))
	}
	return &Scene{
		Width:        800,
		Height:       80,
		Title:        "Genuary - Day13",
		Img:          img,
		FireGradient: gradient,
		Particles:    particles,
	}
}

func (s *Scene) GetDimensions() (int, int) {
	return int(s.Width), int(s.Height)
}

func (s *Scene) Update() error {
	newParticles := make([]*Particle, 0)
	for x := 0; x < 800; x++ {
		newParticles = append(newParticles, NewParticle(float64(x), hScaled-1, s.FireGradient, s.Img))
	}
	for _, particle := range s.Particles {
		particle.Update()
		if particle.LifeSpan > 0 {
			newParticles = append(newParticles, particle)
		}
	}
	s.Particles = newParticles
	return nil
}

func (s *Scene) WriteGif(screen *ebiten.Image) {
	if s.GifWriter == nil {
		s.GifWriter = utils.NewGifWriter("day13.gif", 200)
	}
	img := screen.SubImage(screen.Bounds())
	err := s.GifWriter.RecordGif(img)
	if err != nil {
		fmt.Println(err)
	}
}

func (s *Scene) Draw(screen *ebiten.Image) {
	for _, particle := range s.Particles {
		particle.Draw(screen)
	}
}

func (s *Scene) Layout(oW, oH int) (int, int) {
	return s.GetDimensions()
}
