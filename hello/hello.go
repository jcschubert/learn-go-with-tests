package main

import "fmt"

// Hello provides a greeting to the world
func Hello(name string) string {
	return "Hello, " + name + "."
}

func main() {
	fmt.Println(Hello("Jan"))
}
