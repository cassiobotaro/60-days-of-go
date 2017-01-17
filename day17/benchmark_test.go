package main

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

// only a great number
const size = 100000

func TestSortInts(t *testing.T) {
	// Create some cases to test sort function
	for _, test := range []struct {
		input    []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{[]int{4, 3, 2, 1}, []int{1, 2, 3, 4}},
		{[]int{5, 2, 1, 3}, []int{1, 2, 3, 5}},
	} {
		SortInts(test.input)
		// DeepEqual to compare all values in a list
		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("Expected %v but found %v", test.expected, test.input)
		}
	}
}

func BenchmarkSortInts(b *testing.B) {
	// randomize
	rand.Seed(time.Now().UTC().UnixNano())
	// initialize big array of ints
	bigArray := make([]int, size)
	// with random numbers
	for i := 0; i < size; i++ {
		bigArray[i] = rand.Intn(size)
	}
	// benchmark sort function
	for i := 0; i < b.N; i++ {
		SortInts(bigArray)
	}
}
