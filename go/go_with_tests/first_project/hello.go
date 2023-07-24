package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	englishHello = "Hello, "
	spanishHello = "Hola, "
	frenchHello  = "Bonjour, "
)

func Hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(lang) + name + "!"
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishHello
	case french:
		prefix = frenchHello
	default:
		prefix = englishHello
	}
	return
}

func main() {
	fmt.Println(Hello("world", "English"))
}
