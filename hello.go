package main

import "fmt"

// Hello prints Hello, world.
func Hello(s string) string {
	return "Hello, " + s
}

func main() {
	fmt.Println(Hello("test"))
}
