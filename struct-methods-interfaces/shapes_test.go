package shapes

import "testing"

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
	r := Rectangle{10.0, 10.0}
	got := Area(r)
	want := 100.0
	AssertFloat(t, got, want)
}
