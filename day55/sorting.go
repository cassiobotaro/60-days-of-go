package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age  int
}

type byNameLength []person

// Len is a method needed to implements Interface interface
// represent the slice length
func (s byNameLength) Len() int {
	return len(s)
}

// Swap is a method needed to implements Interface interface
// how elements changes your position
func (s byNameLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less is a method needed to implements Interface interface
// compare method used in sort method
func (s byNameLength) Less(i, j int) bool {
	return len(s[i].Name) < len(s[j].Name)
}

func main() {
	// order strings
	strs := []string{"Felipe", "Cássio", "Gustavo", "César"}
	sort.Strings(strs)
	fmt.Println("Ordered slice: ", strs)

	// order ints
	ints := []int{7, 2, 5, 1, 4, 8, 6, 3}
	sort.Ints(ints)
	fmt.Println("Ordered slice: ", ints)

	// verify if slice is sorted
	fmt.Printf("Are sorted? %t\n", sort.IntsAreSorted(ints))

	// important notes
	// slice have to be sorted
	// always verify if returned position is lower than length of slice
	// and the number of that position is equal the searched number
	x := 4
	// binary search
	c := sort.Search(len(ints), func(i int) bool { return ints[i] >= x })
	if x < len(ints) && ints[c] == x {
		fmt.Printf("found %d at index %d in %v\n", x, c, ints)
	} else {
		fmt.Printf("%d not found in %v\n", x, ints)
	}

	// or

	// The slice must be sorted in ascending order
	// The return value is the index to insert x if x is not present (it could be len(a)).
	i := sort.SearchInts(ints, x)
	if x < len(ints) && ints[i] == x {
		fmt.Printf("found %d at index %d in %v\n", x, c, ints)
	} else {
		fmt.Printf("%d not found in %v\n", x, ints)
	}

	people := []person{
		{"Mari", 24},
		{"Venilton", 60},
		{"Cassio", 26},
	}
	// sorting persons by name length
	sort.Sort(byNameLength(people))
	fmt.Println(people)
	// Interesting links
	// - https://gobyexample.com/sorting
	// - https://gobyexample.com/sorting-by-functions
	// - https://golang.org/pkg/sort/

}
