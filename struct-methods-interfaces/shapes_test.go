package shapes

import (
	"math"
	"testing"
)

func AssertFloat(t *testing.T, got float64, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	r := Rectangle{10.0, 10.0}
	got := Perimeter(r)
	want := 40.0
	AssertFloat(t, got, want)
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		r := Rectangle{10.0, 10.0}
		got := r.Area()
		want := 100.0
		AssertFloat(t, got, want)
	})
	t.Run("circles", func(t *testing.T) {
		c := Circle{10}
		got := c.Area()
		want := math.Pi * 10 * 10
		AssertFloat(t, got, want)
	})
}
