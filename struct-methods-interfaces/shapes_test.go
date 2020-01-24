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

func CheckArea(t *testing.T, shape Shape, want float64) {
	t.Helper()
	got := shape.Area()
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
		CheckArea(t, r, 100.0)
	})
	t.Run("circles", func(t *testing.T) {
		c := Circle{10}
		CheckArea(t, c, math.Pi*10*10)
	})
}
