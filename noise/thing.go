package noise

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mazznoer/colorgrad"
	"github.com/ojrac/opensimplex-go"
	"image"
)

type Thing struct {
	Img   *ebiten.Image
	W     int
	H     int
	Scale float64
	Noise opensimplex.Noise
	Grad  *colorgrad.Gradient
}

func NewThing(x, y, w, h int, scale float64) *Thing {
	grad := colorgrad.Rainbow().Sharp(7, 0.2)
	noise := opensimplex.NewNormalized(934)
	img := image.NewRGBA(image.Rect(x, y, w, h))
	for j := 0; j < h; j++ {
		for i := 0; i < w; i++ {
			t := noise.Eval2(float64(i)*scale, float64(j)*scale)
			img.Set(i, j, grad.At(t))
		}
	}
	return &Thing{
		Img:   ebiten.NewImageFromImage(img),
		Grad:  &grad,
		Noise: noise,
		Scale: scale,
		W:     w,
		H:     h,
	}
}

func (t *Thing) update(time float64) {
	factor := t.Scale * time
	for j := 0; j < t.H; j++ {
		for i := 0; i < t.W; i++ {
			c := t.Noise.Eval2(float64(i)*factor, float64(j)*factor)
			t.Img.Set(i, j, t.Grad.At(c))
		}
	}
}

func (t *Thing) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(t.Img, op)
}
