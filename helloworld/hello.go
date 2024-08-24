package main

import "fmt"

const korean = "Korean"
const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const koreanHelloPrefix = "Ahn nyeong, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case korean:
		prefix = koreanHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println((Hello("DK", "")))
}
