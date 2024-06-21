package models

import (
	"fmt"
	"math"
)

type Load struct {
	Driver     int    // driver number
	LoadNumber string // Load number
	Point      Point
	// Followings will be calculated whenever the Load is introduced
	DistBE float64 // distance between begin and end
	DistB0 float64 // distance between begin and (0.0, 0.0) coordinate
	DistE0 float64 // distance between end and (0.0, 0.0) coordinate
}

type Point struct {
	Visited bool

	Bx, By float64 // begin x, begin y coordinates
	Ex, Ey float64 // end x, end y coordinates
}

func (l *Load) String() string {
	return fmt.Sprintf("ln: %s, B(%f, %f), E(%f, %f)\ndistB0: %f, distE0: %f, distBE: %f",
		l.LoadNumber, l.Point.Bx, l.Point.By, l.Point.Ex, l.Point.Ey, l.DistB0, l.DistE0, l.DistBE)
}

func (l *Load) CalculateDistances() {
	l.DistB0 = math.Sqrt(l.Point.Bx*l.Point.Bx + l.Point.By*l.Point.By)
	l.DistE0 = math.Sqrt(l.Point.Ex*l.Point.Ex + l.Point.Ey*l.Point.Ey)
	l.DistBE = math.Sqrt(l.Point.Bx*l.Point.Bx + l.Point.Ey*l.Point.Ey)

}
