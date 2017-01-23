package main

import "fmt"

func main() {
	// Known in other languages as Dicts, Hashes or associative arrays
	// Maps are a key, value struct

	// map with int as key and int as value
	// sample1 is only declared but not initialized
	var sample1 map[int]int
	// decale and initialize
	// map with string as key and value
	var sample2 = map[string]string{}
	// map with rune as key and bool as value
	sample3 := make(map[rune]bool)
	// initialize with some value
	var sample4 = map[string]int{"z": 1, "y": 3}

	// fill sample1
	// don't forget to initialize
	sample1 = make(map[int]int)
	sample1[1] = 1
	sample1[2] = 1
	sample1[3] = 1

	// print content
	fmt.Printf("sample1 = %+v\n", sample1)

	// length of map
	fmt.Printf("len(sample1) = %+v\n", len(sample1))

	// fill sample2
	sample2["verb"] = "do"
	// retrieve value
	value, ok := sample2["verb"]
	if ok {
		fmt.Println("Retrieved value with success!")
	}
	// ok is False when value is not found
	value, ok = sample2["wrong"]
	if !ok {
		fmt.Println("Value not found!")
		fmt.Println("When not found the value, ok is False and value is zero value.")
		fmt.Printf("%q\n", value)
	}

	// eliminate dupliated runes
	var set = []rune{}
	for _, element := range []rune{'a', 'b', 'a', 'c', 'a', 'a'} {
		if _, ok = sample3[element]; !ok {
			set = append(set, element)
			sample3[element] = true
		}
	}
	fmt.Printf("set = %+v\n", set)

	// iterate over a string and count leters
	for _, letter := range "some word" {
		sample4[string(letter)]++
	}
	fmt.Println(sample4)

	// combining maps and array
	sample5 := map[string][]int{}
	sample5["even"] = []int{2, 4, 6, 8}
	sample5["odd"] = []int{1, 3, 5, 7}
	fmt.Printf("sample5 = %+v\n", sample5)

	// iterate over maps
	// iterate over a map returns key and value
	for key, value := range sample1 {
		fmt.Println("key: ", key, "value: ", value)
	}

}
