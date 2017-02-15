package main

// Run: go test -v

// IsOdd is a function to test if a number is odd
func IsOdd(number int) bool {
	return number%2 != 0
}

// Retangle will implement generate interface in test only for tests purposes
type Retangle struct {
	height int
	width  int
}

// area of retangle
func (r Retangle) area() int {
	return r.height * r.width
}
