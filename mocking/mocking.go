package main

import (
	"fmt"
	"io"
	"os"
)

// Countdown counts down and prints messages, waiting
// 1 second between each message
func Countdown(out io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprintf(out, "Go!")
}

func main() {
	Countdown(os.Stdout)
}
