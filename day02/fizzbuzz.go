package main

import (
	"fmt"
	"strconv"
)

/*
FizzBuzz

In this problem, you should display a list from 1 to 100, one on each line, with the following exceptions:
Numbers divisible by 3 should appear as 'Fizz' instead of number;
Numbers divisible by 5 should appear as 'Buzz' instead of number;
Numbers divisible by 3 and 5 should appear as' FizzBuzz 'instead of number'.
*/

// FizzBuzz should have comments or will not be exported
func FizzBuzz(number int) string {
	// switch with no condition
	// case determines the flow
	switch {
	case number%15 == 0:
		return "FizzBuzz"
	// break is not necessary.
	case number%5 == 0:
		return "Buzz"

	case number%3 == 0:
		return "Fizz"
	default:
		// convert an integer into string
		return strconv.Itoa(number)
	}
}
func main() {
	// Go has only one looping construct, the for loop.
	// Unlike other languages like C, Java, or Javascript there are no parentheses surrounding the three components of the for statement and the braces { } are always required.
	for i := 1; i <= 100; i++ {
		fmt.Println(FizzBuzz(i))
	}
}
