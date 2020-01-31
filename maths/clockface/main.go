package main

import(
	"os"
	"time"

	"github.com/jcschubert/learn-go-with-tests/maths"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}