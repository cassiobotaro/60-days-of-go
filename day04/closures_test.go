package main

import "testing"

// TestSequence tests if sequence is generated correctly
func TestSequence(t *testing.T) {
	// initialize the sequence(closure is returned)
	nextInt := intSequence()
	for i := 0; i < 10; i++ {
		// call nextInt n times
		if nextInt() != i {
			t.Errorf("expected %d but %d was obtained", sequence[i], value)
		}
	}
}
