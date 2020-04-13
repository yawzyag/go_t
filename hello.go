package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const ruso = "Ruso"
const prefixEnglish = "Hello, "
const prefixSpanish = "Hola, "
const prefixFrench = "Bonjour, "
const prefixRuso = "Привет, "

// Hello prints Hello, world.
func Hello(s string, languaje string) string {
	if s == "" {
		s = "world"
	}
	return greetingPrefix(languaje) + s
}

func greetingPrefix(languaje string) (printPrefix string) {
	switch languaje {
	case spanish:
		printPrefix = prefixSpanish
	case french:
		printPrefix = prefixFrench
	case ruso:
		printPrefix = prefixRuso
	default:
		printPrefix = prefixEnglish
	}
	return
}

func main() {
	fmt.Println(Hello("test", "Spanish"))
}
