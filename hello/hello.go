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

	prefix := ""

	switch language {
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("Jan", ""))
}
