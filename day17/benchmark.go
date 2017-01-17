package main

import (
	"fmt"
	"sort"
)

// SortInts sort a list of numbers
func SortInts(nms []int) {
	sort.Ints(nms)
}

func main() {
	// a list of 6 numbers
	numbers := []int{5, 7, 3, 2, 9, 1}
	// sort numbers
	SortInts(numbers)
	// print sorted numbers
	fmt.Printf("numbers = %+v\n", numbers)
}
