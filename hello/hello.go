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

	return getPrefix(language) + name
}

func getPrefix(language string) string {
	switch language {
	case french:
		return frenchPrefix
	case spanish:
		return spanishPrefix
	default:
		return englishPrefix
	}
}

func main() {
	fmt.Println(Hello("Jan", ""))
}
