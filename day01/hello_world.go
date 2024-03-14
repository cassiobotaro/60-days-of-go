// Every program should have a main package
// Declare a main package (a package is a way to group functions, and it's made up of all the files in the same directory).
package main

// fmt package implements formatted I/O with functions analogous to C's printf and scanf. Pronounce is "fumpt"
// Import the popular fmt package, which contains functions for formatting text, including printing to the console. This package is one of the standard library packages you got when you installed Go.
import (
	"fmt"
	"rsc.io/quote"
)

// Every program should have a main function too, and will be the first executed function.
// Implement a main function to print a message to the console. A main function executes by default when you run the main package.
func main() {
	// Println prints "Hello World!!" on the standard output and also break a line.
	fmt.Println("Hello! Newton is trying to learn Go, and He is willing to help you understand it")
	fmt.Println(quote.Glass())
}
