package shapes

import "math"

// Shape is any geometric form with an area.
type Shape interface {
	Area() float64
}

// Rectangle describes a rectangle by its width and height.
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calculates the area of the rectangle r
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle describes a circle by its radius.
type Circle struct {
	Radius float64
}

// Area calculates the area of the circle c
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle is a Shape describe by the length of its base and its height.
type Triangle struct {
	Base   float64
	Height float64
}

// Area calculates the area of the triangle t
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

// Perimeter calculates the perimeter of the rectangle r
func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

// Area calculates the area of the rectangle r
func Area(r Rectangle) float64 {
	return r.Width * r.Height
}
