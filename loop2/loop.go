package loop2

import (
	"fmt"
	"genuary2022/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/ojrac/opensimplex-go"
	"image/color"
	"math"
	"math/rand"
)

type Particle struct {
	Pt         utils.Point
	OriginalPt utils.Point
	CenterPt   utils.Point
	Img        *ebiten.Image
	Color      color.Color
}

func (p *Particle) Draw(screen *ebiten.Image) {
	x := p.Pt.X
	y := p.Pt.Y
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-p.CenterPt.X, -p.CenterPt.Y)
	op.GeoM.Translate(x, y)
	r, g, b, a := p.Color.RGBA()
	op.ColorM.Scale(float64(r), float64(g), float64(b), float64(a))
	screen.DrawImage(p.Img, op)
}

func NewParticle(x, y float64, img *ebiten.Image) *Particle {
	rect := img.Bounds()
	cX := math.Abs(float64(rect.Max.X-rect.Min.X)) / 2
	cY := math.Abs(float64(rect.Max.Y-rect.Min.Y)) / 2
	return &Particle{
		Pt:         utils.Pt(x, y),
		OriginalPt: utils.Pt(x, y),
		CenterPt:   utils.Pt(cX, cY),
		Color:      color.RGBA{R: 0xff, G: 0xff, A: 0x0f},
		Img:        img,
	}
}

type Game struct {
	Width     int
	Height    int
	Scale     int
	Img       *ebiten.Image
	Time      float64
	Radius    float64
	M         float64
	Rad       float64
	NPeriod   float64
	Noises    []opensimplex.Noise
	Particles []*Particle
}

func heartX(radius, p float64) float64 {
	return radius / 15.0 * 16.0 * math.Pow(math.Sin(p), 3)
}

func heartY(radius, p float64) float64 {
	return radius / 15.0 * (-13.0*math.Cos(p) + 5*math.Cos(2*p) + 2*math.Cos(3*p) + math.Cos(4*p))
}

func (g *Game) DrawLine() {
	particles := make([]*Particle, 0)
	for i := 0; i < g.Width; i++ {
		p := float64(i) / g.M
		x := float64(i)
		y := float64(g.Height)/2.0 + g.Noises[0].Eval3(g.Rad*math.Cos(2*math.Pi*g.NPeriod*p), g.Rad*math.Sin(2*math.Pi*g.NPeriod*p), 0.0)*250.0
		particles = append(particles, NewParticle(x, y, g.Img))
	}
	g.Particles = particles
}

func (g *Game) DrawPoints() {
	particles := make([]*Particle, 0)
	for i := 0; i < 2000; i++ {
		a := float64(i) * math.Pi / 180.0
		r := g.Radius * rand.Float64()
		x := float64(g.Width/2.0) + r*math.Cos(a)
		y := float64(g.Height/2.0) + r*math.Sin(a)

		particles = append(particles, NewParticle(x, y, g.Img))
	}
	g.Particles = particles

}

func NewGame(w, h, radius float64, s int) *Game {
	img := ebiten.NewImage(s, s)
	img.Fill(color.White)
	noises := make([]opensimplex.Noise, 0)
	noises = append(noises, opensimplex.New(994))
	noises = append(noises, opensimplex.New(673))
	particles := make([]*Particle, 0)
	game := &Game{
		Width:     int(w),
		Height:    int(h),
		Noises:    noises,
		Scale:     s,
		Radius:    radius,
		NPeriod:   1.5,
		Rad:       0.5,
		M:         1000,
		Particles: particles,
		Img:       img,
	}
	game.DrawPoints()
	return game
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

func (g *Game) AnimateParticles() {
	center := utils.Pt(float64(g.Width/2), float64(g.Height/2))
	for i, particle := range g.Particles {
		p := float64(i) / g.M
		pt := particle.OriginalPt
		sub := pt.Sub(center)
		nx := g.Rad * math.Cos(2*math.Pi*(g.NPeriod*p-g.Time))
		ny := g.Rad * math.Sin(2*math.Pi*(g.NPeriod*p-g.Time))
		ratio := (g.Radius - sub.Length()) / g.Radius
		dx := g.Noises[1].Eval3(nx, ny, 2.0*p) * 200.0 * ratio
		dy := g.Noises[0].Eval3(nx, ny, 2.0*p) * 200.0 * ratio
		particle.Pt = utils.Pt(pt.X+dx, pt.Y+dy)
	}
}

func (g *Game) Update() error {
	g.Time += 0.01
	g.AnimateParticles()
	return nil
}
