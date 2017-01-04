package main

import "testing"

func TestFibonacci(t *testing.T) {
	// first 10 numbers of fibonacci sequence
	sequence := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	// index to iterate over the sequence
	i := 0
	// receive a channel and iterate over it
	for value := range Fib(10) {
		if sequence[i] != value {
			t.Errorf("expected %d but %d was obtained", sequence[i], value)
		}
		i++
	}
}
