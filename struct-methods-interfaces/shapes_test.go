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
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10.0, 10.0}, 100.0},
		{Circle{10}, math.Pi * 10.0 * 10.0},
		{Triangle{12, 6}, 36.0},
	}

	for _, test := range areaTests {
		got := test.shape.Area()
		if got != test.want {
			t.Errorf("got %.2f want %.2f", got, test.want)
		}
	}
}
