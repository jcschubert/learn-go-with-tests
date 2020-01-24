package shapes

// Rectangle describes a rectangle by its width and height.
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter calculates the perimeter of the rectangle r
func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

// Area calculates the area of the rectangle r
func Area(r Rectangle) float64 {
	return r.Width * r.Height
}
