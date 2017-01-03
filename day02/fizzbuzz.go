package main

import "fmt"

/*
FizzBuzz

In this problem, you should display a list from 1 to 100, one on each line, with the following exceptions:
Numbers divisible by 3 should appear as 'Fizz' instead of number;
Numbers divisible by 5 should appear as 'Buzz' instead of number;
Numbers divisible by 3 and 5 should appear as' FizzBuzz 'instead of number'.
*/

// FizzBuzz should have comments or will not be exported
func FizzBuzz(cnt chan int, msg chan string) {
	for {
		i := <-cnt
		// switch with no condition
		// case determines the flow
		switch {
		case i%15 == 0:
			msg <- "FizzBuzz"
		// break is not necessary.
		case i%3 == 0:
			msg <- "Fizz"
		case i%5 == 0:
			msg <- "Buzz"
		default:
			// convert an integer into string
			msg <- fmt.Sprintf("%d", i)
		}
	}
}
func main() {
	// Go has only one looping construct, the for loop.
	// Unlike other languages like C, Java, or Javascript there are no parentheses surrounding the three components of the for statement and the braces { } are always required.
	count := make(chan int)
	message := make(chan string)
	go FizzBuzz(count, message)
	for i := 1; i < 101; i++ {
		count <- i
		fmt.Println(<-message)
	}
}
