package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const prefixEnglish = "Hello, "
const prefixSpanish = "Hola, "
const prefixFrench = "Bonjour, "

// Hello prints Hello, world.
func Hello(s string, languaje string) string {
	printPrefix := prefixEnglish
	if languaje == spanish {
		printPrefix = prefixSpanish
	}
	if languaje == french {
		printPrefix = prefixFrench
	}
	if s == "" {
		s = "world"
	}
	return printPrefix + s
}

func main() {
	fmt.Println(Hello("test", "Spanish"))
}
