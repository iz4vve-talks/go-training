package main

import "fmt"

func main() {
	// var declaration
	var x int = 6

	// var declaration with type inference
	var y = 1.5

	// short-hand declaration (only works in functions)
	z := "Hello, World!"

	// string interpolation and printing
	fmt.Printf("x = %d, y = %.2f, z = %s\n", x, y, z)

	// the fmt package provides printing functions (Printf, Println, Print)
}
