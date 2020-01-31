package clockface

import (
	"testing"
	"time"
)

func TestSecondHandAt30Seconds(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

	want := Point{X: clockCentreX, Y: clockCentreY + secondHandLength}
	got := SecondHand(tm)

	if !pointWithinTolerance(got, want, 0.00001) {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}

