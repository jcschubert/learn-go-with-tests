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

// Sleep allows us to track the amount of calls to it
func (s *SpySleeper) Sleep() {
	s.NumCalls++
}

const write = "write"
const sleep = "sleep"

// CountdownOperationsSpy allows us to spy on both Sleep() and what is written
type CountdownOperationsSpy struct {
	Calls []string
}

// Sleep tracks sleep calls (implements Sleeper)
func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// Write tracks write calls (implements io.Writer)
func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
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
