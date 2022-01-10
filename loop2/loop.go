package loop2

import (
	"fmt"
	"genuary2022/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/ojrac/opensimplex-go"
	"image/color"
	"math"
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
	Time      float64
	Radius    float64
	M         float64
	Rad       float64
	NPeriod   float64
	Noises    []opensimplex.Noise
	Particles []*Particle
}

func NewGame(w, h, radius float64, s int) *Game {
	img := ebiten.NewImage(s, s)
	img.Fill(color.White)
	noises := make([]opensimplex.Noise, 0)
	noises = append(noises, opensimplex.New(994))
	noises = append(noises, opensimplex.New(673))
	particles := make([]*Particle, 0)
	m := 1500.0
	rad := 0.5
	NPeriod := 5.0
	for i := 0; i < int(w); i++ {
		p := float64(i) / m
		x := float64(i)
		y := h/2 + noises[0].Eval3(rad*math.Cos(2*math.Pi*NPeriod*p), rad*math.Sin(2*math.Pi*NPeriod*p), 0.0)*250.0
		particles = append(particles, NewParticle(x, y, img))
	}
	return &Game{
		Width:     int(w),
		Height:    int(h),
		Noises:    noises,
		Scale:     s,
		Radius:    radius,
		NPeriod:   NPeriod,
		Rad:       rad,
		M:         m,
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
	g.Time += 0.01
	h := float64(g.Height / 2.0)
	for i, particle := range g.Particles {
		p := float64(i) / g.M
		nx := g.Rad * math.Cos(2*math.Pi*(g.NPeriod*p-g.Time))
		ny := g.Rad * math.Sin(2*math.Pi*(g.NPeriod*p-g.Time))
		dx := g.Noises[1].Eval3(nx, ny, 4.0*p) * 100.0
		dy := g.Noises[0].Eval3(nx, ny, 4.0*p) * 200.0
		pt := particle.OriginalPt
		particle.Pt = utils.Pt(pt.X+dx, h+dy)
	}
	return nil
}
