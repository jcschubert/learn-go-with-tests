package main

import "fmt"

const englishPrefix = "Hello, "

// Hello provides a greeting to the world
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "Spanish" {
		return "Hola, " + name
	}
	return englishPrefix + name
}

func main() {
	fmt.Println(Hello("Jan", ""))
}
