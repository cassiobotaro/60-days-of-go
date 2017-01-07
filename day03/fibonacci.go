package main

import "fmt"

// Similar code in python:
/*
def fib(n):
	a, b = 0, 1
	for i in range(n):
		a, b = b, a + b
		yield a

for x in fib(10):
	print(x)

# Iterate fib sequence is easy, but next calls and try/exception can be confuse.
*/

// Fib generates fibonacci sequence for n numbers
func Fib(n int) chan int {
	// Create a channel
	c := make(chan int)
	// raise an async go function that generates the fib sequence
	go func() {
		a, b := 0, 1
		for i := 0; i < n; i++ {
			a, b = b, a+b
			// send next value to channel
			c <- a
		}
		// when finished, close the channel.
		// Indicates that all values have already been generated,
		// then stop the iteration
		close(c)
	}()
	// return the channel to function caller
	return c
}

func main() {
	// Iterate over the values received by channel
	// every time a new value has arrived, a loop is executed
	for x := range Fib(10) {
		fmt.Println(x)
	}
}
