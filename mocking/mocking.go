package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countDownStart = 3

// Sleeper allows us to construct a mock for time.Sleep
type Sleeper interface {
	Sleep()
}

// DefaultSleeper wraps time.Sleep in an attached method
type DefaultSleeper struct {
}

// Sleep just calls sleep
func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// SpySleeper tracks the number of calls to it.
type SpySleeper struct {
	NumCalls int
}

//
func (s *SpySleeper) Sleep() {
	s.NumCalls++
}

// Countdown counts down and prints messages, waiting
// 1 second between each message
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countDownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprintf(out, finalWord)
}

func main() {
	s := &DefaultSleeper{}
	Countdown(os.Stdout, s)
}
