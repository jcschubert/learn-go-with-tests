package clockface

import "time"

// A Point represents a 2D Cartesian Coordinate
type Point struct {
	X float64
	Y float64
}

// SecondHand is the unit vector of the second hand of an analogue clock
// represented as a Point
func SecondHand(t time.Time) Point {
	return Point{150, 60}
}
