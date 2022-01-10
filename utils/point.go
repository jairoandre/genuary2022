package utils

type Point struct {
	X float64
	Y float64
}

func Pt(X, Y float64) Point {
	return Point{
		X: X,
		Y: Y,
	}
}

func (p Point) Add(other Point) Point {
	return Point{
		X: p.X + other.X,
		Y: p.Y + other.Y,
	}
}

func (p Point) Sub(other Point) Point {
	return Point{
		X: p.X - other.X,
		Y: p.Y - other.Y,
	}
}

func (p Point) Mul(k float64) Point {
	return Point{
		X: p.X * k,
		Y: p.Y * k,
	}

}
