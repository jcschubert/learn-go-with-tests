package main

import (
	"fmt"
	"io"
	"os"
)

const finalWord := "Go!"
const countDownStart := 3

// Countdown counts down and prints messages, waiting
// 1 second between each message
func Countdown(out io.Writer) {
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprintf(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
