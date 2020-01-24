package shapes

// Perimeter calculates the perimeter of the rectangle defined
// by width and height
func Perimeter(width, height float64) float64 {
	return 2 * (width + height)
}

// Area calculates the area of the rectangle defined
// by width and height
func Area(width, height float64) float64 {
	return width * height
}
