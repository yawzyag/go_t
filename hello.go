package main

import "fmt"

const prefixEnglish = "Hello, "

// Hello prints Hello, world.
func Hello(s string) string {
	if s == "" {
		s = "world"
	}
	return prefixEnglish + s
}

func main() {
	fmt.Println(Hello("test"))
}
