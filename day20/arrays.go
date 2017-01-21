package main

import "fmt"

func main() {
	// declare an array of strings
	var sample1 [5]string
	fmt.Printf("sample1 = %+v\n", sample1)
	// zero values are used to fill the array
	fmt.Printf("sample1[0] = %q\n", sample1[0])

	// fill some position
	sample1[2] = "some string"
	// get position 2
	fmt.Println("sample1[2] -> ", sample1[2])

	// size and capacity
	fmt.Printf("len(sample1) = %+v\n", len(sample1))
	fmt.Printf("cap(sample1) = %+v\n", cap(sample1))

	// Even if not initialized to its full capacity, the array is filled with zero values
	// another way to declare and initialize an array
	var sample2 = [4]int{1, 2}
	fmt.Printf("len(sample2) = %+v\n", len(sample2))
	fmt.Printf("cap(sample2) = %+v\n", cap(sample2))
	fmt.Printf("sample2 = %+v\n", sample2)

	// declare an struct Person
	type Person struct {
		name string
		age  int
	}
	// declare and initialize an array of person
	sample3 := [5]Person{}
	// fill the first person
	sample3[0].name = "Cássio"
	sample3[0].age = 26
	// print index 0
	fmt.Printf("sample3[0] = %+v\n", sample3[0])

	// multi dimensional array
	var matrix [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = i + j
		}
	}
	fmt.Println("matrix: ", matrix)

}
