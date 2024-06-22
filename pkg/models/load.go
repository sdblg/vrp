package models

import (
	"fmt"
)

type Load struct {
	Driver      int      // driver number
	LoadNumbers []string // Load numbers shows which load joined with other loads
	Point       Point
	Joined      bool

	// Followings will be calculated whenever the Load is introduced
	DistBE float64 // load distance between begin and end
	DistB0 float64 // load distance between begin and (0.0, 0.0) coordinate
	DistE0 float64 // load distance between end and (0.0, 0.0) coordinate
	Cost   float64 // load's cost not included driver's cost
}

func (l *Load) String() string {
	return fmt.Sprintf("ln: %s, B(%f, %f), E(%f, %f)\ndistB0: %f, distE0: %f, distBE: %f",
		l.LoadNumbers, l.Point.Bx, l.Point.By, l.Point.Ex, l.Point.Ey, l.DistB0, l.DistE0, l.DistBE)
}

func (l *Load) InitializeDistances() {
	l.DistB0 = l.Point.CalculateDistB0()
	l.DistBE = l.Point.CalculateDistBE()
	l.DistE0 = l.Point.CalculateDistE0()

	l.Cost = l.DistB0 + l.DistBE + l.DistE0
}

// if commit is set as true then the loads will be joined together
// otherwise, respond the joined load cost
func (l *Load) Join(t *Load, commit bool) (cost float64) {
	pointJoin := &Point{
		Bx: l.Point.Ex,
		By: l.Point.Ey,

		Ex: t.Point.Bx,
		Ey: t.Point.By,
	}
	jbe := pointJoin.CalculateDistBE()
	cost = l.Cost - l.DistE0 + jbe + t.Cost - t.DistB0
	if commit { // then update
		l.LoadNumbers = append(l.LoadNumbers, t.LoadNumbers...)
		l.Cost = cost
		l.Point.Ex = t.Point.Ex
		l.Point.Ey = t.Point.Ey
		l.DistE0 = l.Point.CalculateDistE0()
		l.DistBE = l.DistBE + jbe + t.DistBE

		t.Joined = true
	}
	return cost
}

func (l *Load) GetCost() float64 {
	return l.DistB0 + l.DistBE + l.DistE0
}
