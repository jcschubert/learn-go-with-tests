package main

import "fmt"

const englishPrefix = "Hello, "

// Hello provides a greeting to the world
func Hello(name string) string {
	return englishPrefix + name
}

func main() {
	fmt.Println(Hello("Jan"))
}
