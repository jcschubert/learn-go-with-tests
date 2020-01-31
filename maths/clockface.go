package clockface

import (
	"fmt"
	"io"
	"math"
	"time"
)

// A Point represents a 2D Cartesian Coordinate
type Point struct {
	X float64
	Y float64
}

// Clock constants
const (
	secondHandLength = 90
	minuteHandLength = 80
	clockCentreX     = 150
	clockCentreY     = 150
)

// SVG snippets
const (
	svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
	<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
	<svg xmlns="http://www.w3.org/2000/svg"
    	width="100%"
    	height="100%"
    	viewBox="0 0 300 300"
		version="2.0">`
	secondHand = `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`
	minuteHand = `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`
	bezel      = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`
	svgEnd     = `</svg>`
)

func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	SecondHand(w, t)
	MinuteHand(w, t)
	io.WriteString(w, svgEnd)
}

func SecondHand(w io.Writer, t time.Time) {
	p := makeHand(secondHandPoint(t), secondHandLength)
	fmt.Fprintf(w, secondHand, p.X, p.Y)
}

func MinuteHand(w io.Writer, t time.Time) {
	p := makeHand(minuteHandPoint(t), minuteHandLength)
	fmt.Fprintf(w, minuteHand, p.X, p.Y)
}

func makeHand(p Point, length float64) Point {
	p = Point{p.X * length, p.Y * length}                     // scale 
	p = Point{p.X, -p.Y}                                      // flip
	return Point{p.X + clockCentreX, p.Y + clockCentreY}	  // translate
}

func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}

func secondsInRadians(t time.Time) float64 {
	return float64(t.Second()) * (math.Pi / 30)
}

func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / 60) +
		(math.Pi / (30 / float64(t.Minute())))
}

func hoursInRadians(t time.Time) float64 {
	return (math.Pi / (6 / float64(t.Hour() % 12))) 
}
