package loop

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
	Time      float64
	Radius    float64
	Noise     opensimplex.Noise
	Particles []*Particle
}

func NewGame(w, h, radius float64, s int) *Game {
	img := ebiten.NewImage(s, s)
	img.Fill(color.White)
	noise := opensimplex.New(332)
	particles := make([]*Particle, 0)
	center := utils.Pt(w/2, h/2)
	tau := math.Pi / 180.0
	for i := 0; i < 5000; i++ {
		a := rand.Float64() * 360 * tau
		r := rand.Float64() * radius
		nx := r * math.Cos(a)
		ny := r * math.Sin(a)
		pt := center.Add(utils.Pt(nx, ny))
		particles = append(particles, NewParticle(pt.X, pt.Y, img))
	}
	return &Game{
		Width:     int(w),
		Height:    int(h),
		Noise:     noise,
		Scale:     s,
		Radius:    radius,
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
	for _, particle := range g.Particles {
		pt := particle.OriginalPt
		ns := 0.02
		dX := pt.X * ns * math.Cos(g.Time)
		dY := pt.Y * ns * math.Sin(g.Time)
		nX := g.Noise.Eval3(dX, dY, 0.0)
		nY := g.Noise.Eval3(dX, dY, 1.0)
		nPt := utils.Pt(nX, nY).Mul(50.0)
		particle.Pt = pt.Add(nPt)
	}
	return nil
}
