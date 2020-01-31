package clockface

import (
	"math"
	"time"
)

// A Point represents a 2D Cartesian Coordinate
type Point struct {
	X float64
	Y float64
}

// Clock constants
const (
	secondHandLength = 90
	clockCentreX     = 150
	clockCentreY     = 150
)

// SecondHand is the unit vector of the second hand of an analogue clock
// represented as a Point
func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength} // translate
	p = Point{p.X, -p.Y}                                      // flip
	p = Point{p.X + clockCentreX, p.Y + clockCentreY}         // translate
	return p
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}

func secondsInRadians(t time.Time) float64 {
	return float64(t.Second()) * (math.Pi / 30)
}
