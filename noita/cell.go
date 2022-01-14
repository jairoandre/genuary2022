package noita

import (
	"genuary2022/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type CellType uint8

const (
	empty CellType = 0
	sand  CellType = 1
	water CellType = 2
)

type Cell struct {
	Tick     int
	Type     CellType
	Pos      utils.Point
	Img      *ebiten.Image
	LifeTime float64
}

func NewCell(x, y float64, img *ebiten.Image, pType CellType) *Cell {
	return &Cell{
		Type: pType,
		Pos:  utils.Pt(x, y),
		Img:  img,
	}
}

func (p *Cell) Draw(screen *ebiten.Image) {
	var clr color.Color = color.White
	if p.Type == sand {
		clr = color.RGBA{0xff, 0xff, 0x00, 0xff}
	} else if p.Type == water {
		clr = color.RGBA{0x00, 0xaa, 0xff, 0xff}
	}
	op := &ebiten.DrawImageOptions{}
	r, g, b, a := utils.NormalizeColor(clr)
	op.ColorM.Scale(r, g, b, a)
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(p.Pos.X*scale, p.Pos.Y*scale)
	screen.DrawImage(p.Img, op)
}
