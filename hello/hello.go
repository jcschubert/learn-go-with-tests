package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, "

// Hello provides a greeting to the world
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "Spanish" {
		return spanishPrefix + name
	}
	if language == "French" {
		return frenchPrefix + name
	}
	return englishPrefix + name
}

func main() {
	fmt.Println(Hello("Jan", ""))
}
