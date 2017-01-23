package main

import "fmt"

func main() {
	// Slices are like arrays but without a defined size
	// declare a slice of strings
	var sample1 []string
	fmt.Printf("sample1 = %+v\n", sample1)

	// append some items
	sample1 = append(sample1, "apple")
	sample1 = append(sample1, "banana")
	sample1 = append(sample1, "orange")

	fmt.Printf("sample1 = %+v\n", sample1)

	// access by position, just like an array
	fmt.Println("sample1[2] -> ", sample1[2])

	// size and capacity
	fmt.Printf("len(sample1) = %+v\n", len(sample1))
	fmt.Printf("cap(sample1) = %+v\n", cap(sample1))

	// Even if it reaches capacity, it expands
	// append more items
	sample1 = append(sample1, "potato")
	sample1 = append(sample1, "grape")

	// look the new capacity
	fmt.Printf("len(sample1) = %+v\n", len(sample1))
	fmt.Printf("cap(sample1) = %+v\n", cap(sample1))

	// append more them one item at once
	sample1 = append(sample1, "pineapple", "lemon", "pear")

	// initialize with some items
	var sample2 = []int{1, 2, 3}
	fmt.Printf("len(sample2) = %+v\n", len(sample2))
	fmt.Printf("cap(sample2) = %+v\n", cap(sample2))

	// another way to initialize
	// initialize with zero values like arrays
	sample3 := make([]int, 5)
	fmt.Printf("sample3 = %+v\n", sample3)

	// slicing
	// first two items
	// can be assigned to a variable
	fmt.Println(sample1[:2])

	//without first two items
	fmt.Println(sample1[2:])

	// items of index 2 and 3
	fmt.Println(sample1[2:4])

	// declare an struct Person
	type Person struct {
		name string
		age  int
	}
	// declare and initialize an slice of person
	sample4 := []Person{}

	// fill the first person
	sample4 = append(sample4, Person{"Cássio", 26})
	// print index 0
	fmt.Printf("sample4[0] = %+v\n", sample4[0])

	// multi dimensional slice
	matrix := make([][]int, 8)
	for i := 0; i < 8; i++ {
		matrix[i] = make([]int, 4)
		for j := 0; j < 4; j++ {
			matrix[i][j] = i + j
		}
	}
	fmt.Println("matrix: ", matrix)

	// iterate over slices
	// similar to arrays
	for index, value := range sample1 {
		fmt.Println("index ", index, " ->  ", value)
	}
}
