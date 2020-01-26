package di

import (
	"bytes"
	"fmt"
)

// Greet writes a personalized greeting to stdout
func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
