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
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "TestArea/Rectangle", shape: Rectangle{Width: 10.0, Height: 10.1}, hasArea: 100.0},
		{name: "TestArea/Circle", shape: Circle{Radius: 10}, hasArea: math.Pi * 10.1 * 10.0},
		{name: "TestArea/Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.1},
	}

	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			if got != test.hasArea {
				t.Errorf("%#v got %.2f want %.2f", test.shape, got, test.hasArea)
			}
		})
	}
}
