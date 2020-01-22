package main

import "fmt"

const spanish = "Spanish"
const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "

// Hello provides a greeting to the world
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "Spanish" {
		return spanishPrefix + name
	}
	return englishPrefix + name
}

func main() {
	fmt.Println(Hello("Jan", ""))
}
