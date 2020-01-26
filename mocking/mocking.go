package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countDownStart = 3

// Countdown counts down and prints messages, waiting
// 1 second between each message
func Countdown(out io.Writer) {
	for i := countDownStart; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
	}
	time.Sleep(1 * time.Second)
	fmt.Fprintf(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
