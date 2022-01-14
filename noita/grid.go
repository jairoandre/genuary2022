package noita

import "github.com/hajimehoshi/ebiten/v2"

type Grid struct {
	Cells [][]*Cell
	Tick  int
}

func (g *Grid) Get(x, y int) *Cell {
	if y < 0 || y >= len(g.Cells) {
		return nil
	}
	row := g.Cells[y]
	if x < 0 || x >= len(row) {
		return nil
	}
	return row[x]
}

func (g *Grid) UpdateSandCell(x, y int) bool {
	cell := g.Get(x, y)
	if cell != nil && cell.Type == empty {
		cell.Type = sand
		cell.Tick = g.Tick
		return true
	}
	return false
}

func (g *Grid) UpdateSand(x, y int) {
	curr := g.Get(x, y)
	if curr == nil {
		return
	}
	if g.Tick == curr.Tick {
		return
	}
	curr.Tick = g.Tick
	curr.Type = empty
	curr.Type = empty
	if g.UpdateSandCell(x, y+1) {
		return
	}
	if g.UpdateSandCell(x-1, y+1) {
		return
	}
	if g.UpdateSandCell(x+1, y+1) {
		return
	}
	curr.Type = sand
}

func (g *Grid) Update() {
	g.Tick += 1
	for y := 0; y < len(g.Cells); y++ {
		row := g.Cells[y]
		for x := 0; x < len(row); x++ {
			particle := row[x]
			switch particle.Type {
			case sand:
				g.UpdateSand(x, y)
			default:
				// nothing
			}
		}
	}
}

func (g *Grid) Draw(screen *ebiten.Image) {
	for _, row := range g.Cells {
		for _, cell := range row {
			if cell.Type == empty {
				continue
			}
			cell.Draw(screen)
		}
	}
}
