package noita

import (
	"genuary2022/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/mazznoer/colorgrad"
	"image/color"
)

const (
	scale   = 2
	width   = 800
	height  = 600
	wScaled = width / scale
	hScaled = height / scale
)

type Scene struct {
	Title        string
	Img          *ebiten.Image
	Grid         *Grid
	FireGradient colorgrad.Gradient
	GifWriter    *utils.GifWriter
	IsPainting   bool
}

func NewScene() *Scene {
	img := ebiten.NewImage(1, 1)
	img.Fill(color.White)
	gradient := colorgrad.Inferno()
	grid := Grid{
		Cells: make([][]*Cell, 0),
	}
	for y := 0; y < hScaled; y++ {
		row := make([]*Cell, 0)
		for x := 0; x < wScaled; x++ {
			xf64 := float64(x)
			yf64 := float64(y)
			row = append(row, NewCell(xf64, yf64, img, empty))
		}
		grid.Cells = append(grid.Cells, row)
	}
	return &Scene{
		Title:        "Noita Go",
		Img:          img,
		FireGradient: gradient,
		Grid:         &grid,
	}
}

func (s *Scene) GetDimensions() (int, int) {
	return width, height
}

func (s *Scene) Painting(cType CellType) {
	mx, my := ebiten.CursorPosition()
	rx := mx / scale
	ry := my / scale
	if rx > 0 && rx < wScaled && ry > 0 && ry < hScaled {
		for j := -4; j <= 4; j++ {
			for i := -4; i <= 4; i++ {
				cell := s.Grid.Get(rx+i, ry+j)
				if cell == nil {
					continue
				}
				cell.Type = cType
			}
		}
	}
}

func (s *Scene) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		s.IsPainting = true
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		s.IsPainting = false
	}
	if s.IsPainting {
		s.Painting(sand)
	} else {
		if s.Grid.Tick%10 == 0 {
			s.Grid.Cells[0][wScaled/2].Type = sand
		}
		s.Grid.Update()
	}
	return nil
}

func (s *Scene) Draw(screen *ebiten.Image) {
	s.Grid.Draw(screen)
	utils.DebugInfo(screen)
}

func (s *Scene) Layout(oW, oH int) (int, int) {
	return s.GetDimensions()
}
