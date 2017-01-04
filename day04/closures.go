package main

import "fmt"

// Closures are functions that refer to independent (free) variables
// (variables that are used locally, but defined in an enclosing scope).
// In other words, these functions 'remember' the environment in which they were created.

// intSeq returns a function that returns a sequence of integers
func intSequence() func() int {
	// name "i" is binded when closure is created
	i := 0
	// returns a function that don't receive parameters, and returns a number
	return func() int {
		// increase value and return i
		i++
		return i
	}
}

func main() {
	nextInt := intSequence()
	fmt.Println(nextInt()) // When the function is called the value of free variable "i" is 0
	//it will print 1(because of i++)
	fmt.Println(nextInt()) // when called again, the function "remember" the value of "i"
	//it will print 2
	fmt.Println(nextInt()) // again the function "remeber" the value of the "i"
	// finally, it returns 3
	// This function generates integers infinitely
}
