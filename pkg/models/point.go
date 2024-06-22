package models

import "math"

type Point struct {
	Bx, By float64 // begin x, begin y coordinates
	Ex, Ey float64 // end x, end y coordinates
}

func (p *Point) CalculateDistB0() float64 {
	return math.Sqrt(p.Bx*p.Bx + p.By*p.By)
}

func (p *Point) CalculateDistE0() float64 {
	return math.Sqrt(p.Ex*p.Ex + p.Ey*p.Ey)
}

// Distance between Begining and End point
func (p *Point) CalculateDistBE() float64 {
	return math.Sqrt((p.Ex-p.Bx)*(p.Ex-p.Bx) + (p.Ey-p.By)*(p.Ey-p.By))
}
