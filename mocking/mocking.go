package main

import (
	"fmt"
	"io"
	"os"
)

// Countdown counts down and prints messages, waiting
// 1 second between each message
func Countdown(out io.Writer) {
	fmt.Fprintf(out, "3")
}

func main() {
	Countdown(os.Stdout)
}
